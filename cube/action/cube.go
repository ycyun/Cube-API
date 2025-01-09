package action

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	Cube "github.com/ycyun/Cube-API/cube/model"
	"net/http"
)

func InitCube() {
	Cube.Cube()
	Cube.Hosts()
}

func PrintError() {
	fmt.Println(json.MarshalIndent(Cube.Cube(), "1", "2"))
}

// Version godoc
//
//	@Summary		Show Versions of CUBE
//	@Description	CUBE 의 버전을 보여줍니다.
//	@Tags			API, CUBE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	utils.TypeVersion
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/version [get]
func Version(context *gin.Context) {
	dat := Cube.Cube().GetVersion()
	// Print the output
	dat.Debug = gin.IsDebugging()
	context.IndentedJSON(http.StatusOK, dat)
} // @name Version

func UpdateHosts() {
	Cube.UpdateHosts()
}

// GetHosts godoc
//
//	@Summary		Show List of Host
//	@Description	Cube의 Host목록을 보여줍니다.
//	@Tags			API, CUBE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	Cube.TypeHosts
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/cube/hosts [get]
func GetHosts(ctx *gin.Context) {
	ret := Cube.Hosts()
	ctx.IndentedJSON(http.StatusOK, ret)
}
