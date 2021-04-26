package server

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"

	"stash.us.cray.com/rabsw/nnf-ec/internal/common"
	"stash.us.cray.com/rabsw/nnf-ec/internal/logging"
)

type DefaultServerControllerProvider struct{}

func (DefaultServerControllerProvider) NewServerController(opts ServerControllerOptions) ServerControllerApi {
	if opts.Local {
		return &LocalServerController{}
	}
	return NewRemoteServerController(opts)
}

type LocalServerController struct {
	storage []Storage
}

func (c *LocalServerController) Connected() bool { return true }

func (c *LocalServerController) NewStorage(pid uuid.UUID) *Storage {
	c.storage = append(c.storage, Storage{Id: pid, ctrl: c})
	return &c.storage[len(c.storage)-1]
}

func (c *LocalServerController) Delete(s *Storage) error {
	for storageIdx, storage := range c.storage {
		if storage.Id == s.Id {
			c.storage = append(c.storage[:storageIdx], c.storage[storageIdx+1:]...)
			break
		}
	}

	return nil
}

func (c *LocalServerController) GetStatus(s *Storage) StorageStatus {

	// We really shouldn't need to refresh on every GetStatus() call if we're correctly
	// tracking udev add/remove events. There should be a single refresh on launch (or
	// possibily a udev-info call to pull in the initial hardware?)
	if err := c.Discover(nil); err != nil {
		logrus.WithError(err).Errorf("Local Server Controller: Discovery Error")
		return StorageStatus_Error
	}

	// There should always be 1 or more namespces, so zero namespaces mean we are still starting.
	// Once we've recovered the expected number of namespaces (nsExpected > 0) we continue to
	// return a starting status until all namespaces are available.

	// TODO: We should check for an error status somewhere... as this stands we are not
	// ever going to return an error if refresh() fails.
	if (len(s.ns) == 0) ||
		((s.nsExpected > 0) && (len(s.ns) < s.nsExpected)) {
		return StorageStatus_Starting
	}

	if s.nsExpected == len(s.ns) {
		return StorageStatus_Ready
	}

	return StorageStatus_Error
}

func (c *LocalServerController) CreateFileSystem(s *Storage, fs FileSystemApi, mp string) error {
	s.fileSystem = fs

	opts := FileSystemCreateOptions{
		"mountpoint": mp,
	}

	return fs.Create(s.Devices(), opts)
}

func (c *LocalServerController) DeleteFileSystem(s *Storage) error {
	return s.fileSystem.Delete()
}

func (c *LocalServerController) Discover(newStorageFunc func(*Storage)) error {
	nss, err := c.namespaces()
	if err != nil {
		return err
	}

	for _, ns := range nss {
		sns, err := c.newStorageNamespace(ns)

		if errors.Is(err, common.ErrNamespaceMetadata) {
			continue
		} else if err != nil {
			return err
		}

		s := c.findStorage(sns.poolId)
		if s == nil {
			s = c.NewStorage(sns.poolId)

			s.nsExpected = sns.poolTotal

			s.ns = append(s.ns, *sns)

			if newStorageFunc != nil {
				newStorageFunc(s)
			}

		} else {

			// We've identified a pool for this particular namespace
			// Add the namespace to the pool if it's not present.
			s.UpsertStorageNamespace(sns)

		}

	}

	return nil
}

func (c *LocalServerController) findStorage(pid uuid.UUID) *Storage {
	for idx, p := range c.storage {
		if p.Id == pid {
			return &c.storage[idx]
		}
	}

	return nil
}

func (c *LocalServerController) namespaces() ([]string, error) {
	nss, err := c.run("ls -A1 /dev/nvme* | grep -E \"nvme[0-9]+n[0-9]+\"")

	// The above will return an err if zero namespaces exist. In
	// this case, ENOENT is returned and we should instead return
	// a zero length array.
	if exit, ok := err.(*exec.ExitError); ok {
		if syscall.Errno(exit.ExitCode()) == syscall.ENOENT {
			return make([]string, 0), nil
		}
	}

	return strings.Fields(string(nss)), err
}

func (c *LocalServerController) newStorageNamespace(path string) (*StorageNamespace, error) {

	// First we need to identify the NSID for the provided path.
	nsid, err := getNamespaceId(path)
	if err != nil {
		return nil, err
	}

	// Retrieve the namespace GUID
	guidStr, err := c.run(fmt.Sprintf("nvme id-ns %s --namespace-id=%d | awk '/nguid/{printf $3}'", path, nsid))
	if err != nil {
		return nil, err
	}

	id, err := uuid.ParseBytes(guidStr)
	if err != nil {
		return nil, err
	}

	data, err := c.run(fmt.Sprintf("nvme get-feature %s --namespace-id=%d --feature-id=0x7F --raw-binary", path, nsid))
	if err != nil {
		return nil, err
	}

	meta, err := common.DecodeNamespaceMetadata(data[6:])
	if err != nil {
		return nil, err
	}

	return &StorageNamespace{
		id:        id,
		path:      path,
		nsid:      nsid,
		poolId:    meta.Id,
		poolIdx:   int(meta.Index),
		poolTotal: int(meta.Count),
	}, nil
}

func (c *LocalServerController) run(cmd string) ([]byte, error) {
	return logging.Cli.Trace(cmd, func(cmd string) ([]byte, error) {
		return exec.Command("bash", "-c", cmd).Output()
	})
}

// https://elixir.bootlin.com/linux/latest/source/include/uapi/asm-generic/ioctl.h
const (
	_IOC_NONE  = 0x0
	_IOC_WRITE = 0x1
	_IOC_READ  = 0x2

	_IOC_NRBITS   = 8
	_IOC_TYPEBITS = 8
	_IOC_SIZEBITS = 14
	_IOC_DIRBITS  = 2

	_IOC_NRSHIFT   = 0
	_IOC_TYPESHIFT = _IOC_NRSHIFT + _IOC_NRBITS
	_IOC_SIZESHIFT = _IOC_TYPESHIFT + _IOC_TYPEBITS
	_IOC_DIRSHIFT  = _IOC_SIZESHIFT + _IOC_SIZEBITS
)

func _IOC(dir uint, t uint, nr uint, size uint) uint {
	return (dir << _IOC_DIRSHIFT) |
		(t << _IOC_TYPESHIFT) |
		(nr << _IOC_NRSHIFT) |
		(size << _IOC_SIZESHIFT)
}

func _IO(t uint, nr uint) uint { return _IOC(_IOC_NONE, t, nr, 0) }

// https://github.com/linux-nvme/nvme-cli/blob/master/linux/nvme_ioctl.h
var _NVME_IOCTL_ID = func() uint { return _IO('N', 0x40) }()

func getNamespaceId(path string) (int, error) {
	fd, err := unix.Open(path, unix.O_RDONLY, 0)
	if err != nil {
		return -1, err
	}
	defer unix.Close(fd)

	return unix.IoctlRetInt(fd, _NVME_IOCTL_ID)
}
