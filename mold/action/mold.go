package action

import (
	"github.com/gin-gonic/gin"
	Mold "github.com/ycyun/Cube-API/mold/model"
	"github.com/ycyun/Cube-API/utils"
	"net/http"
	"time"
)

// GetStatus godoc
//
//	@Summary		Show Status of MOLD
//	@Description	MOLD의 상태값을 보여줍니다.
//	@Tags			API, Mold, MOLD
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	Mold.TypeMoldStatus
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/mold [get]
func GetStatus(ctx *gin.Context) {

	ctx.IndentedJSON(http.StatusOK, utils.TestReturn{RefreshTime: Mold.Status().RefreshTime, Now: time.Now()})
}

func MonitorStatus() {
	Mold.UpdateStatus()
}
