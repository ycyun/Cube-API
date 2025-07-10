package cube

import (
	"github.com/ycyun/Cube-API/Modules/cube/DiskController"
	"log/slog"
	"reflect"
	"sync"
)

type StructCube struct {
	ID       string                         `json:"id"`
	HostName string                         `json:"hostname"`
	IP       string                         `json:"ip"`
	Disks    *DiskController.DiskController `json:"disks"`
	NicList  []string                       `json:"nic_list"`
}

// Psamples는 StructSamples의 싱글톤 인스턴스로, 지연 초기화되며 스레드 안전성을 위해 sync.Once로 보호됩니다.
var lockCube sync.Once
var Cube *StructCube

func Init() *StructCube {
	if Cube == nil {
		lockCube.Do(
			func() {
				slog.Debug("Create Struct", "type", reflect.TypeOf(Cube))
				Cube = &StructCube{
					Disks:   DiskController.Init(),
					NicList: []string{},
					IP:      "0.0.0.0",
				}
			})
	} else {
		//slog.Debug(fmt.Sprintf("get old ", reflect.TypeOf(Cube), " instance."))
		slog.Debug("Use Old Struct", "type", reflect.TypeOf(Cube))
	}
	return Cube
}

func cubeTest() {
	c := Init()
	c.GetDiskList()
}
