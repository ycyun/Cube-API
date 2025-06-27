package DiskController

import "github.com/ycyun/Cube-API/Modules/cube/DiskController/Disk"

type Interface interface {
	GetDisks() []*Disk.Interface
	GetRBDs() []*Disk.Interface
	GetBLKs() []*Disk.Interface
	UpdateRBDs()
	UpdateBLKs()
}
