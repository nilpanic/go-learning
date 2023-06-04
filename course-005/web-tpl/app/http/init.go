package http

import (
	"github.com/gin-gonic/gin"

	"web-tpl/app"
	"web-tpl/app/http/routes"
)

func NewServer() error {
	// 启动 server
	r := gin.Default()
	//_ = r.SetTrustedProxies(nil)

	// 其它配置
	routes.Reg(r)

	return r.Run(app.Config.HTTP.Listen) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
