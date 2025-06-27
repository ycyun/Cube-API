package watcher

import (
	"github.com/ycyun/Cube-API/utils"
	"log"
	"time"
)

func (worker *Watcher) StatusRegister(fn func()) {

	worker.Handlers = append(worker.Handlers, fn)
}
func (worker *Watcher) Run() {
	// 환경 변수 및 설정 로드
	cfg := utils.LoadConfig()
	worker.running = true
	// 주기적으로 상태 업데이트 실행
	ticker := time.NewTicker(time.Duration(cfg.TaskInterval) * time.Second)
	defer ticker.Stop()

	log.Println("Watcher started...")

	for range ticker.C {
		log.Println("Running background tasks...")
		if worker.running {
			for _, handler := range worker.Handlers {
				go handler()
			}
		}
	}
}

func (worker *Watcher) Stop() {
	worker.running = false
}
