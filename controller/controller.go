package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/ycyun/Cube-API/utils"
	"io"
	"log"
	"net/http"
	"reflect"
	"slices"
	"sync"
	"time"
)

type TypeNeighbor struct {
	IP       string `json:"ip"`
	HostName string `json:"hostname"`
}

type TypeNeighbors struct {
	Neighbors []TypeNeighbor `json:"neighbors"`
}

type TypeController struct {
	Handlers []func()          `json:"handlers"`
	running  bool              `json:"running"`
	Neighbor TypeNeighbors     `json:"neighbors"`
	errors   utils.Errors      `json:"errors"`
	version  utils.TypeVersion `json:"version"`
	Cube     interface{}       `json:"cube"`
	Mold     interface{}       `json:"mold_status"`
	Glue     interface{}       `json:"glue_status"`
} //	@name	TypeController

var lockController sync.Once
var controller *TypeController

func Init() *TypeController {
	if controller == nil {
		lockController.Do(
			func() {
				fmt.Println("Creating ", reflect.TypeOf(controller), " now.")
				controller = &TypeController{}
			})
	} else {
		fmt.Println("get old ", reflect.TypeOf(controller), " instance.")
	}
	return controller
}

func (c *TypeController) StatusRegister(fn func()) {

	c.Handlers = append(c.Handlers, fn)
}

func (c *TypeController) Start() {
	c.running = true
	for c.running {
		for _, handler := range c.Handlers {
			go handler()
		}

		time.Sleep(time.Duration(10000) * time.Millisecond)
	}
}

func (c *TypeController) Stop() {
	c.running = false
}

func (c *TypeController) AddError(err error) {
	serr := err.Error()
	c.errors.Errors = append(c.errors.Errors, utils.Errorlog{Error: serr, Time: time.Now()})
}

func AddError(err error) {
	controller.AddError(err)
}

func (c *TypeController) GetError() utils.Errors {
	return c.errors
}

func (c *TypeController) ClearError() {
	c.errors = utils.Errors{}
}

// Error godoc
//
//	@Summary		Error
//	@Description	Error.
//	@Tags			API, CUBE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	utils.Errorlog
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/err [get]
func (c *TypeController) Error(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, c.GetError())
}

func (c *TypeController) DeleteError(context *gin.Context) {
	c.ClearError()
	context.IndentedJSON(http.StatusOK, c.GetError())
}

// GetNeighbor godoc
//
//	@Summary		GetNeighbor
//	@Description	GetNeighbor.
//	@Tags			API, CUBE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	TypeNeighbors
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/cube/neighbor [get]
func (c *TypeController) GetNeighbor(context *gin.Context) {
	ret := c.Neighbor
	context.IndentedJSON(http.StatusOK, ret)
}

func (c *TypeController) LoadConfig() {

}

// PutNeighbor godoc
//
//	@Summary		PutNeighbor
//	@Description	PutNeighbor.
//	@Tags			API, CUBE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			ip			formData	string	true	"Neighbor IP"
//	@Param			hostname	formData	string	true	"Neighbor Hostname"
//	@Success		200			{object}	TypeNeighbors
//	@Failure		400			{object}	HTTP400BadRequest
//	@Failure		404			{object}	HTTP404NotFound
//	@Failure		500			{object}	HTTP500InternalServerError
//	@Router			/cube/neighbor [post]
//	@Router			/cube/neighbor [put]
func (c *TypeController) PutNeighbor(ctx *gin.Context) {
	var neighbor TypeNeighbor

	neighbor.IP = ctx.PostForm("ip")
	neighbor.HostName = ctx.PostForm("hostname")
	if err := ctx.ShouldBindQuery(&neighbor); err != nil {
		log.Println("err: ", err)
	}
	if err := ctx.ShouldBindUri(&neighbor); err != nil {
		log.Println("err: ", err)
	}
	if err := ctx.ShouldBindJSON(&neighbor); err != nil {
		log.Println("err: ", err)
	}

	log.Printf("user: %+v", neighbor)
	c.Neighbor.Neighbors = append(c.Neighbor.Neighbors, neighbor)
	SaveConfig()
	ctx.IndentedJSON(http.StatusOK, c.Neighbor)
}

// DeleteNeighbor godoc
//
//	@Summary		DeleteNeighbor
//	@Description	DeleteNeighbor.
//	@Tags			API, CUBE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			ip			formData	string	true	"Neighbor IP"
//	@Param			hostname	formData	string	true	"Neighbor Hostname"
//	@Success		200			{object}	TypeNeighbor
//	@Failure		400			{object}	HTTP400BadRequest
//	@Failure		404			{object}	HTTP404NotFound
//	@Failure		500			{object}	HTTP500InternalServerError
//	@Router			/cube/neighbor [delete]
func (c *TypeController) DeleteNeighbor(ctx *gin.Context) {
	var neighbor TypeNeighbor

	neighbor.IP = ctx.PostForm("ip")
	neighbor.HostName = ctx.PostForm("hostname")
	if err := ctx.ShouldBindQuery(&neighbor); err != nil {
		log.Println("err: ", err)
	}
	if err := ctx.ShouldBindUri(&neighbor); err != nil {
		log.Println("err: ", err)
	}
	if err := ctx.ShouldBindJSON(&neighbor); err != nil {
		log.Println("err: ", err)
	}

	log.Printf("neighbor: %+v", neighbor)
	if neighbor.IP != "" && neighbor.HostName != "" {

		for i, n := range c.Neighbor.Neighbors {
			if n.IP == neighbor.IP && n.HostName == neighbor.HostName {
				c.Neighbor.Neighbors = slices.Delete(c.Neighbor.Neighbors, i, i+1)
			}
		}
	} else if neighbor.IP == "" && neighbor.HostName != "" {
		for i, n := range c.Neighbor.Neighbors {
			if n.HostName == neighbor.HostName {
				c.Neighbor.Neighbors = slices.Delete(c.Neighbor.Neighbors, i, i+1)
			}
		}
	} else if neighbor.IP != "" && neighbor.HostName == "" {
		for i, n := range c.Neighbor.Neighbors {
			if n.IP == neighbor.IP {
				c.Neighbor.Neighbors = slices.Delete(c.Neighbor.Neighbors, i, i+1)
			}
		}
	} else {
		for i, _ := range c.Neighbor.Neighbors {
			c.Neighbor.Neighbors = slices.Delete(c.Neighbor.Neighbors, i, i+1)
		}
	}
	SaveConfig()
	ctx.IndentedJSON(http.StatusOK, c.Neighbor)
}

// GetNeighborInfo godoc
//
//	@Summary		GetNeighbor
//	@Description	GetNeighbor.
//	@Tags			API, CUBE
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	TypeNeighborInfos
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/cube/neighbor/info [get]
func (c *TypeController) GetNeighborInfo(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, c.UpdateNeighborInfo())
}

func (n *TypeNeighbor) GetFromNeighbor(api string) (map[string]interface{}, int) {
	req, err := http.NewRequest("GET", "http://"+n.IP+":8080/api/"+api, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("User-Agent", "CubeAPI")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력

	bytes, _ := io.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로
	fmt.Println(str)
	var ret map[string]interface{}
	err = json.Unmarshal([]byte(str), &ret)
	if err != nil {
		fmt.Println(err)
		controller.AddError(err)
		ret["err"] = err
		return ret, resp.StatusCode
	}
	return ret, resp.StatusCode
}

type TypeNeighborInfos struct {
	Neighbors map[string]TypeNeighborInfo `json:"neighbors"`
}
type TypeNeighborInfo struct {
	Info map[string]interface{} `json:"info"`
	Code int                    `json:"code"`
}

func (c *TypeController) UpdateNeighborInfo() TypeNeighborInfos {
	ret := TypeNeighborInfos{Neighbors: make(map[string]TypeNeighborInfo)}
	for _, neighbor := range c.Neighbor.Neighbors {
		str, code := neighbor.GetFromNeighbor("v1/cube/neighbor")
		ret.Neighbors[neighbor.HostName] = TypeNeighborInfo{str, code}
	}
	return ret
}
