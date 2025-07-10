package BLK

import "github.com/ycyun/Cube-API/Modules/cube/DiskController/Disk"

type BLK struct {
	*Disk.Base
	Name        string        `json:"name"`
	MajMin      string        `json:"maj:min"`
	Rm          bool          `json:"rm"`
	Size        string        `json:"size"`
	Ro          bool          `json:"ro"`
	Type        string        `json:"type"`
	Mountpoints []interface{} `json:"mountpoints"`
	Children    []BLK         `json:"children"`
}

type LsblkOutput struct {
	BlockDevices []*BLK `json:"blockdevices"`
}
