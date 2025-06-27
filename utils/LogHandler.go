package utils

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"
)

var AbleJsonHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{

	Level: slog.LevelInfo,
	//AddSource: true,
	ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			return slog.Attr{}
		}
		/*
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)
				a.Value = slog.AnyValue(slog.Source{
					Line: source.Line,
					File:     source.File, // + ":" + strconv.Itoa(source.Line),
					Function: source.Function,
				})
			}
		*/
		return a
	},
})
var AbleTextHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{

	Level: slog.LevelInfo,
	//AddSource: true,
	ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			return slog.Attr{}
		}
		return a
	},
})

var AbleHandler = AbleTextHandler

var Logger = slog.New(AbleHandler)

// HandleError this logs the function name as well.
func HandleError(err error) (b bool) {
	if err != nil {
		// notice that we're using 1, so it will actually log the where
		// the Error happened, 0 = this function, we don't want that.
		// pc, filename, line, _ := runtime.Caller(1)
		_, filename, line, _ := runtime.Caller(1)

		//log.Printf("[Error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
		//slog.Error(fmt.Sprintf("err %v", err))
		slog.Error(fmt.Sprintf("%s:%d %v", filename, line, err))
		// the format should be "Modules/api/routes.go:13:23: undefined: c"
		b = true
	}
	return
}

type Errors struct {
	Errors []Errorlog `json:"errors"`
}
type Errorlog struct {
	Error string    `json:"error" format:"string"`
	Time  time.Time `json:"refresh_time" format:"time"`
} // @name Errorlog

func LogInit() *slog.Logger {

	slog.SetDefault(Logger) // default 설정. logger 대신 slog로 로그 찍어도 logger랑 똑같은 기능을 함.
	slog.SetLogLoggerLevel(slog.LevelInfo)
	return Logger
}
