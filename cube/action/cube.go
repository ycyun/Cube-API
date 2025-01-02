package action

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	Cube "github.com/ycyun/Cube-API/cube/model"
)

func StatusRegister(fn func()) {
	Cube.Cube().StatusRegister(fn)
}

func Start() {
	Cube.Cube().Start()
}

func Stop() {
	Cube.Cube().Stop()
}

func AddError(err error) {
	Cube.Cube().AddError(err)
}
func Error(context *gin.Context) {
	Cube.Cube().Error(context)
}

func ClearError(context *gin.Context) {
	Cube.Cube().ClearError(context)
}

func Init() *Cube.TypeCUBE {
	return Cube.Cube()
}

func PrintError() {
	fmt.Println(json.MarshalIndent(Cube.Cube(), "1", "2"))
}

func Version(context *gin.Context) {
	Cube.Cube().Version(context)
}
