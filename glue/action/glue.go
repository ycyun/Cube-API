package action

import (
	"github.com/gin-gonic/gin"
	Glue "github.com/ycyun/Cube-API/glue/model"
	"net/http"
)

// GetGlueStatus godoc
//
//	@Summary		Show Status of GLUE
//	@Description	GLUE의 상태값을 보여줍니다.
//	@Tags			API, Glue, GLUE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	Glue.TypeGlueStatus
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/glue [get]
func GetGlueStatus(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, Glue.Status())
}

func MonitorGlueStatus() {
	Glue.UpdateStatus()
}
