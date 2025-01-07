package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	Cube "github.com/ycyun/Cube-API/cube/action"
	"os/exec"
	"reflect"
	"sync"
	"time"
)

// TypeGlueVersion
// @Description Glue의 버전
type TypeGlueVersion struct {
	Mon interface {
	} `json:"mon"`
	Mgr interface {
	} `json:"mgr"`
	Osd interface {
	} `json:"osd"`
	RbdMirror interface {
	} `json:"rbd-mirror"`
	Rgw interface {
	} `json:"rgw"`
	Overall interface {
	} `json:"overall"`
} //@name TypeGlueVersion

var lockGlueVersion sync.Once

var _glueVersion *TypeGlueVersion

func Version() *TypeGlueVersion {
	if _glueVersion == nil {
		lockGlueVersion.Do(
			func() {
				if gin.IsDebugging() {
					fmt.Println("Creating ", reflect.TypeOf(_glueVersion), " now.")
				}
				_glueVersion = &TypeGlueVersion{}
			})
	} else {
		if gin.IsDebugging() {
			fmt.Println("using old glueversion instance now.")
		}
	}
	return _glueVersion
}

func UpdateVersion() *TypeGlueVersion {
	if gin.Mode() == gin.ReleaseMode {
		cmd := exec.Command("ceph", "versions")
		stdout, _ := cmd.CombinedOutput()

		if err := json.Unmarshal(stdout, &_glueVersion); err != nil {
			Cube.AddError(err)
		}
	} else {
		// Print the output
		versions := []byte("{\n    \"mon\": {\n        \"ceph version Glue-Diplo-4.0.0 (5dd24139a1eada541a3bc16b6941c5dde975e26d) reef (stable)\": 3\n    },\n    \"mgr\": {\n        \"ceph version Glue-Diplo-4.0.0 (5dd24139a1eada541a3bc16b6941c5dde975e26d) reef (stable)\": 2\n    },\n    \"osd\": {\n        \"ceph version Glue-Diplo-4.0.0 (5dd24139a1eada541a3bc16b6941c5dde975e26d) reef (stable)\": 19\n    },\n    \"rbd-mirror\": {\n        \"ceph version Glue-Diplo-4.0.0 (5dd24139a1eada541a3bc16b6941c5dde975e26d) reef (stable)\": 2\n    },\n    \"rgw\": {\n        \"ceph version Glue-Diplo-4.0.0 (5dd24139a1eada541a3bc16b6941c5dde975e26d) reef (stable)\": 1\n    },\n    \"overall\": {\n        \"ceph version Glue-Diplo-4.0.0 (5dd24139a1eada541a3bc16b6941c5dde975e26d) reef (stable)\": 27\n    }\n}")
		if err := json.Unmarshal(versions, &_glueVersion); err != nil {
			Cube.AddError(err)
		}
	}
	_glueStatus.RefreshTime = time.Now()
	return _glueVersion
}
