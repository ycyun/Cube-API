package watcher

import (
	"github.com/ycyun/Cube-API/utils"
	"log/slog"
	"reflect"
	"sync"
)

type Watcher struct {
	Handlers []func() `json:"handlers"`
	running  bool
}

var lockWatcher sync.Once // Worker의 락
var Pwatcher *Watcher     // Worker의 포인터

func Init(conf *utils.StructConfig) *Watcher {
	if Pwatcher == nil {
		lockWatcher.Do(
			func() {
				////fmt.Println("Creating ", reflect.TypeOf(Pwatcher), "with config", conf, " now.")
				slog.Debug("Create Struct", "type", reflect.TypeOf(Pwatcher), "config", conf)
				Pwatcher = &Watcher{}
			})
	} else {
		//fmt.Println("Creating ", reflect.TypeOf(Pwatcher), "with config", conf, " now.")
		slog.Debug("Use Old Struct", "type", reflect.TypeOf(Pwatcher), "config", conf)

	}
	return Pwatcher
}
