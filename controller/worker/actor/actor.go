package actor

import (
	"github.com/ycyun/Cube-API/controller/worker/jobqueue"
	"github.com/ycyun/Cube-API/utils"
	"log"
	"os/exec"
	"time"
)

func (worker *Actor) Run() {
	cfg := utils.LoadConfig()
	worker.running = true
	// 주기적으로 상태 업데이트 실행
	ticker := time.NewTicker(time.Duration(cfg.TaskInterval) * time.Second)
	defer ticker.Stop()

	log.Println("Actor started...")

	for range ticker.C {
		job := worker.jq.Pop()
		payload := job.Payload
		switch payload.Type {
		case "cmd":
			exec.Command(payload.Cmd, payload.Args...)
			log.Println("Shell script task running")
		case "rest":
			log.Println("Rest task running")
		}
		log.Println("Running background tasks...")
	}
}

func (worker *Actor) Stop() {
	worker.running = false
}
func (worker *Actor) AddJob(job *jobqueue.Job) {
	worker.jq.Push(job)
	worker.running = false
}
