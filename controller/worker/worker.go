package worker

import (
	"fmt"
	"github.com/ycyun/Cube-API/Modules/config"
	"log"
	"reflect"
	"sync"
	"time"
)

type APIWorker struct {
	Handlers []func()             `json:"handlers"`
	Config   *config.StructConfig `json:"cfg"`
	running  bool
}

var lockAPI sync.Once
var Pworker *APIWorker

func Init(conf *config.StructConfig) *APIWorker {
	if Pworker == nil {
		lockAPI.Do(
			func() {
				fmt.Println("Creating ", reflect.TypeOf(Pworker), "with config", conf, " now.")
				Pworker = &APIWorker{
					Config: conf,
				}
			})
	} else {
		fmt.Println("get old ", reflect.TypeOf(Pworker), " instance.")
	}
	return Pworker
}

func (worker *APIWorker) StatusRegister(fn func()) {

	worker.Handlers = append(worker.Handlers, fn)
}
func (worker *APIWorker) Run() {
	// 환경 변수 및 설정 로드
	cfg := config.Load()
	worker.running = true
	// 주기적으로 상태 업데이트 실행
	ticker := time.NewTicker(time.Duration(cfg.TaskInterval) * time.Second)
	defer ticker.Stop()

	log.Println("APIWorker started...")

	for range ticker.C {
		log.Println("Running background tasks...")
		if worker.running {
			for _, handler := range worker.Handlers {
				go handler()
			}
		}
	}
}

func (worker *APIWorker) Stop() {
	worker.running = false
}
