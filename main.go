package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	Cube "github.com/ycyun/Cube-API/cube/action"
	"github.com/ycyun/Cube-API/docs"
	Glue "github.com/ycyun/Cube-API/glue/action"
	Mold "github.com/ycyun/Cube-API/mold/action"
	"log"
	"time"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
func main() {
	cube := Cube.Init()
	Cube.StatusRegister(Glue.MonitorGlueStatus)
	Cube.StatusRegister(Mold.MonitorStatus)

	cube.StatusRegister(errorMaker)
	go Cube.Start()

	docs.SwaggerInfo.Title = "Cube API"
	docs.SwaggerInfo.Description = "This is a Cube-API server."
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = ".swagger.io"

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies(nil)
	if err != nil {
		Cube.AddError(err)
	}

	r.Use(gin.Logger())

	// Recovery 미들웨어는 panic이 발생하면 500 에러를 씁니다.
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		glue := v1.Group("/glue")
		{
			glue.GET("", Glue.GetGlueStatus)
		}
		mold := v1.Group("/mold")
		{
			mold.GET("", Mold.GetStatus)
		}
		v1.Any("/version", Cube.Version)
		v1.GET("/err", Cube.Error)
		v1.DELETE("/err", Cube.ClearError)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = r.Run(":8080")
	if err != nil {
		Cube.AddError(err)
	}
	Cube.Stop()
	Cube.PrintError()
	fmt.Println("end")
}

func errorMaker() {
	Cube.AddError(errors.New(time.Now().String()))
}
