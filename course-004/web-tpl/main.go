package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"

	"web-tpl/app"
)

func main() {
	// 解析项目根目录参数
	homePath := flag.String("prjHome", "", "项目的根目录路径")
	flag.Parse()

	if *homePath == "" {
		_, f, _, ok := runtime.Caller(0)
		if !ok {
			panic("尝试获取文件路径失败！")
		}

		*homePath = filepath.Dir(f)
	}

	err := app.Init(*homePath)
	if err != nil {
		panic(err)
	}

	// 启动 server
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	fmt.Println(app.Config.HTTP.Enable)

	err = r.Run(app.Config.HTTP.Listen) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
