package jobqueue

import (
	"github.com/ycyun/Cube-API/utils"
	"log/slog"
	"reflect"
	"sync"
)

type _job struct {
	Type string
	Cmd  string
	Args []string
}
type Job struct {
	ID      int
	Payload _job
}

type JobQueue struct {
	Jobs []*Job `json:"jobs"`
}

var lockQueue sync.Once // Worker의 락
var Pqueue *JobQueue    // Worker의 포인터

func Init(conf *utils.StructConfig) *JobQueue {
	if Pqueue == nil {
		lockQueue.Do(
			func() {
				//fmt.Println("Creating ", reflect.TypeOf(Pqueue), "with config", conf, " now.")
				slog.Debug("Create Struct", "type", reflect.TypeOf(Pqueue), "config", conf)
				Pqueue = &JobQueue{}
			})
	} else {
		//fmt.Println("Creating ", reflect.TypeOf(Pqueue), "with config", conf, " now.")
		slog.Debug("Use Old Struct", "type", reflect.TypeOf(Pqueue), "config", conf)

	}
	return Pqueue
}
