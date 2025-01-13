package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	C "github.com/ycyun/Cube-API/controller"
	Cube "github.com/ycyun/Cube-API/cube/action"
	Dashboard "github.com/ycyun/Cube-API/dashboard/action"
	"github.com/ycyun/Cube-API/docs"
	Glue "github.com/ycyun/Cube-API/glue/action"
	Mold "github.com/ycyun/Cube-API/mold/action"
	PCS "github.com/ycyun/Cube-API/pcs/action"
	UTILS "github.com/ycyun/Cube-API/utils"
	"log"
	"time"
)

//	@title			Cube API
//	@version		1.0
//	@description	This is a Cube-API server.
//	@termsOfService	https://ablecloud.io/

//	@contact.name	API Support
//	@contact.url	https://www.ablecloud.io/support
//	@contact.email	ycyun@ablecloud.io

//	@license.name	Apache 2.0
//	@license.url	https://www.apache.org/licenses/LICENSE-2.0.html

//	@ssshost						10.211.55.11:8080
//	@BasePath					/api/v1
//	@Schemes					http https
//	@securityDefinitions.basic	None

// @externalDocs.description	ABLECLOUD
// @externalDocs.url			https://www.ablecloud.io
func main() {
	// 시간대 설정
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		panic(err)
	}
	// Set the timezone for the current process
	time.Local = location

	c := C.Init()
	C.LoadConfig()
	Cube.InitCube()
	c.StatusRegister(Mold.MonitorStatus)
	c.StatusRegister(Glue.Monitor)
	c.StatusRegister(Dashboard.Monitor)
	c.StatusRegister(PCS.Monitor)
	c.StatusRegister(Cube.UpdateHosts)
	c.StatusRegister(C.SaveConfig)

	go c.Start()
	APIPort := "8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = UTILS.GetLocalIP().String() + ":" + APIPort
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := gin.Default()
	//gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)
	r.ForwardedByClientIP = true
	err = r.SetTrustedProxies(nil)
	if err != nil {
		c.AddError(err)
	}

	r.Use(gin.Logger())

	// Recovery 미들웨어는 panic이 발생하면 500 에러를 씁니다.
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		cube := v1.Group("/cube")
		{
			cube.GET("/hosts", Cube.GetHosts)
			cube.GET("/test", Cube.GetHosts)
			cube.GET("/neighbor", c.GetNeighbor)
			cube.GET("/neighbor/info", c.GetNeighborInfo)
			cube.POST("/neighbor", c.PutNeighbor)
			cube.PUT("/neighbor", c.PutNeighbor)
			cube.DELETE("/neighbor", c.DeleteNeighbor)
		}
		glue := v1.Group("/glue")
		{
			glue.GET("/", Glue.GetGlueStatus)
			glue.GET("/auth", Glue.GetGlueAuth)
			glue.GET("/auth/:username", Glue.GetGlueAuth)
			glue.GET("/auths", Glue.GetGlueAuths)
		}
		mold := v1.Group("/mold")
		{
			mold.GET("", Mold.GetStatus)
			mold.GET("/ccvm", Mold.GetCCVMInfo)
		}
		pcs := v1.Group("/pcs")
		{
			pcs.GET("", PCS.GetStatus)
			pcs.GET("/resources", PCS.GetResource)
		}
		dashboard := v1.Group("/dashboard")
		{
			dashboard.GET("", Dashboard.GetStatus)

		}
		v1.Any("/version", Cube.Version)
		v1.GET("/err", c.Error)
		v1.DELETE("/err", c.DeleteError)
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	err = r.Run(":" + APIPort)
	if err != nil {
		c.AddError(err)
	}
	c.Stop()
	Cube.PrintError()
	fmt.Println("end")
}

func errorMaker() {
	c := C.Init()
	c.AddError(errors.New(time.Now().String()))
}
