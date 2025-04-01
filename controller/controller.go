package controller

import (
	"github.com/ycyun/Cube-API/Modules/config"
	"github.com/ycyun/Cube-API/controller/router"
	"github.com/ycyun/Cube-API/controller/worker"
	"sync"
)

type TypeController struct {
	Config *config.StructConfig `json:"cfg"`
	A      *router.APIRouter
	W      *worker.APIWorker
}

var lockController sync.Once
var controller *TypeController

func Init() *TypeController {
	lockController.Do(func() {
		conf := config.Load()
		controller = &TypeController{
			Config: conf,
			A:      router.Init(conf),
			W:      worker.Init(conf),
		}
	})
	return controller
}
