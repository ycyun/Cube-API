package controller

import (
	"github.com/ycyun/Cube-API/Modules/cube"
	"github.com/ycyun/Cube-API/controller/router"
	"github.com/ycyun/Cube-API/controller/worker"
	"github.com/ycyun/Cube-API/utils"
	"sync"
)

type TypeController struct {
	Config *utils.StructConfig `json:"cfg"`
	A      *router.APIRouter
	W      *worker.APIWorker
}

var lockController sync.Once
var controller *TypeController

func Init() *TypeController {
	lockController.Do(func() {
		conf := utils.LoadConfig()
		controller = &TypeController{
			Config: conf,
			A:      router.Init(conf),
			W:      worker.Init(conf),
		}
		cube.Init()
	})
	return controller
}
