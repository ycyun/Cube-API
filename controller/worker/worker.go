package worker

import (
	"github.com/ycyun/Cube-API/Base"
	"github.com/ycyun/Cube-API/utils"
	"log"
	"log/slog"
	"reflect"
	"sync"
	"time"
)

type APIWorker struct {
	Handlers []Base.Interface    `json:"handlers"`
	Config   *utils.StructConfig `json:"cfg"`
	running  bool
}

var lockWorker sync.Once // Worker의 락
var Pworker *APIWorker   // Worker의 포인터

func Init(conf *utils.StructConfig) *APIWorker {
	if Pworker == nil {
		lockWorker.Do(
			func() {
				//fmt.Println("Creating ", reflect.TypeOf(Pworker), "with config", conf, " now.")
				slog.Debug("Create Struct", "type", reflect.TypeOf(Pworker), "config", conf)
				Pworker = &APIWorker{
					Config: conf,
				}
			})
	} else {
		//fmt.Println("Creating ", reflect.TypeOf(Pworker), "with config", conf, " now.")
		slog.Debug("Use Old Struct", "type", reflect.TypeOf(Pworker), "config", conf)
	}
	return Pworker
}

func (worker *APIWorker) StatusRegister(fn Base.Interface) {

	worker.Handlers = append(worker.Handlers, fn)
}
func (worker *APIWorker) Run() {
	// 환경 변수 및 설정 로드
	cfg := utils.LoadConfig()
	worker.running = true
	// 주기적으로 상태 업데이트 실행
	ticker := time.NewTicker(cfg.TaskInterval * time.Second)
	defer ticker.Stop()

	slog.Debug("APIWorker started with ticker interval", cfg.TaskInterval)
	//
	//for range ticker.C {
	//	log.Println("Running background tasks...")
	//	if worker.running {
	//		for _, handler := range worker.Handlers {
	//			go handler.Update()
	//		}
	//	}
	//}

	for worker.running {
		log.Println("Running background tasks...")
		for _, handler := range worker.Handlers {
			go handler.Update()

		}
		time.Sleep(cfg.TaskInterval)
	}

	log.Println("APIWorker stopped...")
}

func (worker *APIWorker) Stop() {
	worker.running = false

}
