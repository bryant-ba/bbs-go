package admin

import (
	"os"
	"runtime"

	"github.com/kataras/iris/v12"
	"github.com/mlogclub/simple/mvc"
)

type CommonController struct {
	Ctx iris.Context
}

func (c *CommonController) GetSysteminfo() *mvc.JsonResult {
	hostname, _ := os.Hostname()
	return mvc.NewEmptyRspBuilder().
		Put("os", runtime.GOOS).
		Put("arch", runtime.GOARCH).
		Put("numCpu", runtime.NumCPU()).
		Put("goversion", runtime.Version()).
		Put("hostname", hostname).
		JsonResult()
}
