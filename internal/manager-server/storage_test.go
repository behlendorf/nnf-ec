package server

import (
	"testing"

	"github.com/google/uuid"
)

func _TestCreateZfsFilesystem(t *testing.T) {
	// Warning, this assume the existence of some NVMe devices

	fsConfig, _ := loadConfig()
	fsCtrl := NewFileSystemController(fsConfig)

	pid := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	provider := DefaultServerControllerProvider{}
	ctrl := provider.NewServerController(ServerControllerOptions{Local: true})
	pool := ctrl.NewStorage(pid)

	if pool == nil {
		t.Fatalf("Could not allocate storage pool %s", pid)
	}

	fs := fsCtrl.NewFileSystem("ZFS")
	if fs == nil {
		t.Fatalf("Could not allocate ZFS file system")
	}

	if err := pool.CreateFileSystem(fs, "/mnt/zfs"); err != nil {
		t.Errorf("Create File System Failed %s", err)
	}
}
