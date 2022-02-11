// 创建 微服务
package server

import (
	"github.com/beego/beego/v2/core/logs"

	"go-micro.dev/v4"
)

func Start() {
	service := micro.NewService(micro.Name("MicroServer"))
	service.Init()
	service.Run()
	logs.Info("Init Server Service Success")
}
