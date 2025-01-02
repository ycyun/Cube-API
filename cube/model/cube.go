package action

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

type errorlog struct {
	Error string    `json:"error" format:"string"`
	Time  time.Time `json:"refresh_time" format:"time"`
} // @name errorlog

type TypeCUBE struct {
	Handlers []func()
	running  bool
	errors   []errorlog
}

var lockCUBE sync.Once
var cube *TypeCUBE

func Cube() *TypeCUBE {
	if cube == nil {
		lockCUBE.Do(
			func() {
				fmt.Println("Creating glue instance now.")
				cube = &TypeCUBE{
					Handlers: []func(){},
					running:  false,
					errors:   []errorlog{},
				}
			})
	} else {
		fmt.Println("glue instance already created.")
	}

	return cube
}

type TypeCUBEVersion struct {
	Version string `json:"version,omitempty"`
	Debug   bool   `json:"debug,omitempty"`
}

// Version godoc
//
//	@Summary		Show Versions of API
//	@Description	API 의 버전을 보여줍니다.
//	@Tags			API
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Success		200	{object}	TypeCUBEVersion
//	@Failure		400	{object}	HTTP400BadRequest
//	@Failure		404	{object}	HTTP404NotFound
//	@Failure		500	{object}	HTTP500InternalServerError
//	@Router			/version [get]
func (c *TypeCUBE) Version(context *gin.Context) {
	dat := TypeCUBEVersion{Version: "1.0.0", Debug: true}
	// Print the output
	dat.Debug = gin.IsDebugging()
	context.IndentedJSON(http.StatusOK, dat)
} // @name Version

func (c *TypeCUBE) StatusRegister(fn func()) {

	c.Handlers = append(c.Handlers, fn)
}

func (c *TypeCUBE) Start() {
	c.running = true
	for c.running {
		for _, handler := range c.Handlers {
			go handler()
		}

		time.Sleep(time.Duration(10000) * time.Millisecond)
	}
}

func (c *TypeCUBE) Stop() {
	c.running = false
}

func (c *TypeCUBE) AddError(err error) {
	serr := err.Error()
	c.errors = append(c.errors, errorlog{Error: serr, Time: time.Now()})
}
func (c *TypeCUBE) Error(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, c.errors)
}

func (c *TypeCUBE) GetError() []errorlog {
	return c.errors
}

func (c *TypeCUBE) ClearError(context *gin.Context) {
	c.errors = []errorlog{}
	context.IndentedJSON(http.StatusOK, c.errors)
}
