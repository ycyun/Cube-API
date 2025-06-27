package DiskController

import (
	"fmt"
	"github.com/ycyun/Cube-API/Modules/cube/DiskController/BLK"
	"github.com/ycyun/Cube-API/Modules/cube/DiskController/Disk"
	"github.com/ycyun/Cube-API/controller/worker"
	"log/slog"
	"reflect"
	"sync"
)

type DiskController struct {
	Disks  []Disk.Interface `json:"disks"`
	RBDs   []Disk.Interface `json:"RBDs"`
	BLKs   []*BLK.BLK       `json:"BLKs"`
	HBAs   []*Disk.Base
	iSCSIs []*Disk.Base
	ID     string `json:"ID"`
} //@name DiskController

func (s *DiskController) GetID() string {
	return s.ID
}

func (s *DiskController) SetID(id string) bool {
	s.ID = id
	return true
}

func (s *DiskController) Update() {
	//slog.Info("DiskController Update")
	s.UpdateRBDs()
	//slog.Info("sRBDs = ", "s", s, s.RBDs)
	tmpList := make([]Disk.Interface, 0)
	s.Disks = append(tmpList, s.RBDs...)
}

// PDC DiskController 싱글톤 인스턴스로, 지연 초기화되며 스레드 안전성을 위해 sync.Once로 보호됩니다.
var lockSample sync.Once
var PDC *DiskController

func Init() *DiskController {
	if PDC == nil {
		lockSample.Do(
			func() {
				fmt.Println("Creating ", reflect.TypeOf(PDC), " now.")
				PDC = &DiskController{
					Disks: make([]Disk.Interface, 0),
					RBDs:  make([]Disk.Interface, 0),
					BLKs:  make([]*BLK.BLK, 0),
				}
				worker.Pworker.StatusRegister(PDC)
			})
	} else {
		//fmt.Println("Creating ", reflect.TypeOf(PDC), "with config", conf, " now.")
		slog.Debug("Use Old Struct", "type", reflect.TypeOf(PDC))
	}
	return PDC
}
