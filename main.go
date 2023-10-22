package main

import (
	"InvertedCow/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

func newApp(engine *gin.Engine, config *config.AppConfig) *http.Server {
	srv := &http.Server{
		Addr:    config.Port,
		Handler: engine,
	}
	return srv
}

func main() {
	//获取参数
	path, _ := os.Getwd()
	path = strings.ReplaceAll(path, "\\", "/")
	path = path + "/conf/config.ini"

	//加载配置
	conf, err := config.InitSetting(path)
	if err != nil {
		fmt.Println("加载配置文件出错")
		return
	}

	if err != nil {
		log.Println(err)
	}

	//注册路由
	srv, err := initApp(conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println(err)
	}
}
