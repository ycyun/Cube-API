package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ycyun/Cube-API/Modules/route" // APIRouter 핸들러 임포트
	"github.com/ycyun/Cube-API/utils"
	"log/slog"
	"reflect"
	"sync"
)

type APIRouter struct {
	Config *utils.StructConfig `json:"cfg"`
}

var lockRouter sync.Once //Router 생성에 대한 lock
var PRouter *APIRouter   //Router의 포인터

func Init(conf *utils.StructConfig) *APIRouter {
	if PRouter == nil {
		lockRouter.Do(
			func() {
				//fmt.Println("Creating ", reflect.TypeOf(PRouter), "with config", conf, " now.")
				slog.Debug("Create Struct", "type", reflect.TypeOf(PRouter), "config", conf)
				PRouter = &APIRouter{
					Config: conf,
				}
			})
	} else {
		//fmt.Println("Creating ", reflect.TypeOf(PRouter), "with config", conf, " now.")
		slog.Debug("Use Old Struct", "type", reflect.TypeOf(PRouter), "config", conf)
	}
	return PRouter
}

func (*APIRouter) Run() {
	// 환경 변수 및 설정 로드
	var err error
	cfg := utils.LoadConfig()

	// Gin 엔진 생성
	router := gin.Default()
	router.Use(gin.Logger())
	// Recovery 미들웨어는 panic이 발생하면 500 에러를 씁니다.
	router.Use(gin.Recovery())

	//gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)
	router.ForwardedByClientIP = true
	err = router.SetTrustedProxies(nil)

	// APIRouter 엔드포인트 등록
	route.RegisterRoutes(router)

	// 서버 실행
	slog.Info("Starting APIRouter server", "port", cfg.ServerPort)
	if err = router.Run(":" + cfg.ServerPort); err != nil {
		slog.Error("Failed to start server", "error", err)
		//c.AddError(err)
	}
}
