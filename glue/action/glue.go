package action

import (
	"github.com/gin-gonic/gin"
	Glue "github.com/ycyun/Cube-API/glue/model"
	"log"
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

// GetGlueAuth godoc
//
//	@Summary		Show Status of GLUE
//	@Description	GLUE의 상태값을 보여줍니다.
//	@Tags			API, Glue, GLUE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	Glue.TypeAuth
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/glue [get]
func GetGlueAuth(ctx *gin.Context) {
	var user Glue.User
	if err := ctx.ShouldBindQuery(&user); err != nil {
		log.Println("err: ", err)
	}
	if err := ctx.ShouldBindUri(&user); err != nil {
		log.Println("err: ", err)
	}
	log.Printf("user: %+v", user)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		log.Println("err: ", err)
	}
	ret := Glue.GetAuth(user)
	if ret != nil {
		ctx.IndentedJSON(http.StatusOK, ret)
	} else {
		ctx.AbortWithStatusJSON(http.StatusNotFound, Glue.TypeAuth{Entity: user.Username, Key: ""})
	}
}

// GetGlueAuths godoc
//
//	@Summary		Show Status of GLUE
//	@Description	GLUE의 상태값을 보여줍니다.
//	@Tags			API, Glue, GLUE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	Glue.TypeAuth
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/glue [get]
func GetGlueAuths(ctx *gin.Context) {
	ret := Glue.GetAuths()
	ctx.IndentedJSON(http.StatusOK, ret)
}

func Monitor() {
	Glue.UpdateStatus()
	Glue.UpdateHealth()
	Glue.UpdateAuths()
}
func MonitorGlueStatus() {
	Glue.UpdateStatus()
}

func MonitorGlueHealthDetail() {
	Glue.UpdateHealth()
}
