package cube

import (
	"github.com/ycyun/Cube-API/Modules/cube/DiskController/Disk"
)

func (c *StructCube) GetDiskList() []Disk.Interface {
	//slog.Info("dDisks: ", c.Disks)
	return c.Disks.Disks
}

func (c *StructCube) GetNicList() []string {
	return c.NicList
}

func (c *StructCube) FormatDisk(fstype string, disks []Disk.Base) (bool, error) {
	for _, disk := range disks {
		if disk.Type == "disk" {
			disk.Type = fstype
		}
	}
	return true, nil
}
