package action

import (
	"github.com/gin-gonic/gin"
	Dashboard "github.com/ycyun/Cube-API/dashboard/model"
	"net/http"
)

// UpdateStatus godoc
//
//	@Summary		Show StorageCenterClusterStatus of GLUE
//	@Description	GLUE의 상태값을 보여줍니다.
//	@Tags			API, Glue, GLUE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	Glue.TypeGlueStatus
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/glue [get]
func UpdateStatus(ctx *gin.Context) {
	SCC := Dashboard.StorageCenterClusterUpdateStatus()
	ret := map[string]interface{}{
		"StorageCenterCluster": SCC,
	}
	ctx.IndentedJSON(http.StatusOK, ret)
}

func Monitor() {
	Dashboard.StorageCenterClusterUpdateStatus()
}
func MonitorDashboard() {
	Dashboard.StorageCenterClusterUpdateStatus()
}
