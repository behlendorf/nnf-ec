/*
 * Copyright 2020, 2021, 2022 Hewlett Packard Enterprise Development LP
 * Other additional copyright holders may be indicated within.
 *
 * The entirety of this work is licensed under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 *
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package server

import (
	"fmt"

	"github.com/NearNodeFlash/nnf-ec/pkg/var_handler"
)

// FileSystemXfs establishes an XFS file system on an underlying LVM logical volume.
type FileSystemXfs struct {
	FileSystemLvm

	MkfsMount FileSystemOemMkfsMount
}

func init() {
	FileSystemRegistry.RegisterFileSystem(&FileSystemXfs{})
}

func (*FileSystemXfs) New(oem FileSystemOem) (FileSystemApi, error) {
	return &FileSystemXfs{
		FileSystemLvm: FileSystemLvm{
			FileSystem: FileSystem{name: oem.Name},
			CmdArgs:    oem.LvmCmd,
		},
		MkfsMount: oem.MkfsMount,
	}, nil
}

func (*FileSystemXfs) IsType(oem FileSystemOem) bool { return oem.Type == "xfs" }
func (*FileSystemXfs) IsMockable() bool              { return false }

func (*FileSystemXfs) Type() string   { return "xfs" }
func (f *FileSystemXfs) Name() string { return f.name }

func (f *FileSystemXfs) Create(devices []string, opts FileSystemOptions) (err error) {
	if err := f.FileSystemLvm.Create(devices, opts); err != nil {
		return err
	}

	varHandler := var_handler.NewVarHandler(map[string]string{
		"$DEVICE": f.FileSystemLvm.devPath(),
	})
	mkfsArgs := varHandler.ReplaceAll(f.MkfsMount.Mkfs)

	if _, err := f.run(fmt.Sprintf("mkfs.xfs %s", mkfsArgs)); err != nil {
		return err
	}

	return setFileSystemPermissions(f, opts)
}

func (f *FileSystemXfs) Mount(mountpoint string) error {
	return f.mount(f.devPath(), mountpoint, "", f.MkfsMount.Mount)
}
