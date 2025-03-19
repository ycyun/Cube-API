package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ycyun/Cube-API/Modules/utils"
	"net/http"
	"os"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1")
	{

		v1.GET("/neighbor", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusOK, "hello")
		})
		v1.GET("/err", func(ctx *gin.Context) {
			_, err := os.Open("test123")
			utils.HandleError(err)
			ctx.IndentedJSON(http.StatusOK, "hello")
		})
	}
	//{
	//	v1.GET("/neighbor", c.GetNeighbor)
	//	v1.GET("/neighbor/info", c.GetNeighborInfo)
	//	v1.POST("/neighbor", c.PutNeighbor)
	//	v1.PUT("/neighbor", c.PutNeighbor)
	//	v1.DELETE("/neighbor", c.DeleteNeighbor)
	//	cube := v1.Group("/cube")
	//	{
	//		cube.GET("/hosts", Cube.Hosts.Get)
	//		cube.GET("/test", Cube.Hosts.Get)
	//		cube.GET("/nics", Cube.NICs.Get)
	//		cube.GET("/disk", Cube.Disks.Get)
	//	}
	//	glue := v1.Group("/glue")
	//	{
	//		glue.GET("/", Glue.GetGlueStatus)
	//		glue.GET("/auth", Glue.GetGlueAuth)
	//		glue.GET("/auth/:username", Glue.GetGlueAuth)
	//		glue.GET("/auths", Glue.GetGlueAuths)
	//	}
	//	mold := v1.Group("/mold")
	//	{
	//		mold.GET("", Mold.GetStatus)
	//		mold.GET("/ccvm", Mold.GetCCVMInfo)
	//	}
	//	pcs := v1.Group("/pcs")
	//	{
	//		pcs.GET("", PCS.GetStatus)
	//		pcs.GET("/resources", PCS.GetResource)
	//	}
	//	dashboard := v1.Group("/dashboard")
	//	{
	//		dashboard.GET("", Dashboard.GetStatus)
	//
	//	}
	//	//v1.Any("/version", Cube.Version)
	//	v1.GET("/err", c.Error)
	//	v1.DELETE("/err", c.DeleteError)
	//	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}
}
