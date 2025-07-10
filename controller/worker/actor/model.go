package actor

import (
	"github.com/ycyun/Cube-API/controller/worker/jobqueue"
	"github.com/ycyun/Cube-API/utils"
	"log/slog"
	"reflect"
	"sync"
)

type Actor struct {
	running bool
	jq      *jobqueue.JobQueue
}

var lockActor sync.Once // Worker의 락
var Pactor *Actor       // Worker의 포인터

func Init(conf *utils.StructConfig) *Actor {
	if Pactor == nil {
		lockActor.Do(
			func() {
				//fmt.Println("Creating ", reflect.TypeOf(Pactor), "with config", conf, " now.")
				slog.Debug("Create Struct", "type", reflect.TypeOf(Pactor), "config", conf)
				Pactor = &Actor{}
			})
	} else {
		//fmt.Println("Creating ", reflect.TypeOf(Pactor), "with config", conf, " now.")
		slog.Debug("Use Old Struct", "type", reflect.TypeOf(Pactor), "config", conf)
	}
	return Pactor
}
