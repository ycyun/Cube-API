package model

import (
	"github.com/ycyun/Cube-API/utils"
)

type gluePools struct {
	Pools []string `json:"pools"`
} // @name GluePools

type Auth struct {
	Token       string `json:"token"`
	Username    string `json:"username"`
	Permissions struct {
		Cephfs            []string `json:"cephfs"`
		ConfigOpt         []string `json:"config-opt"`
		DashboardSettings []string `json:"dashboard-settings"`
		Grafana           []string `json:"grafana"`
		Hosts             []string `json:"hosts"`
		Iscsi             []string `json:"iscsi"`
		Log               []string `json:"log"`
		Manager           []string `json:"manager"`
		Monitor           []string `json:"monitor"`
		NfsGanesha        []string `json:"nfs-ganesha"`
		Osd               []string `json:"osd"`
		Pool              []string `json:"pool"`
		Prometheus        []string `json:"prometheus"`
		RbdImage          []string `json:"rbd-image"`
		RbdMirroring      []string `json:"rbd-mirroring"`
		Rgw               []string `json:"rgw"`
		User              []string `json:"user"`
	} `json:"permissions"`
	PwdExpirationDate interface{} `json:"pwdExpirationDate"`
	Sso               bool        `json:"sso"`
	PwdUpdateRequired bool        `json:"pwdUpdateRequired"`
} // @name Auth

type glueError struct {
	utils.HTTPError
	Detail    string `json:"detail"`
	Code      string `json:"code"`
	Component string `json:"component"`
} // @name GlueError

type glueFS struct {
	Name string `json:"name"`
} // @name GlueFS

type FSList struct {
	glueFS []glueFS `json:"glueFS"`
} // @name FSList

type glueFSInfo struct {
	MonAddrs                  []string `json:"mon_addrs"`
	PendingSubvolumeDeletions int      `json:"pending_subvolume_deletions"`
	Pools                     struct {
		Data []struct {
			Avail int64  `json:"avail"`
			Name  string `json:"name"`
			Used  int    `json:"used"`
		} `json:"data"`
		Metadata []struct {
			Avail int64  `json:"avail"`
			Name  string `json:"name"`
			Used  int    `json:"used"`
		} `json:"metadata"`
	} `json:"pools"`
	UsedSize int `json:"used_size"`
} // @name GlueFSInfo
