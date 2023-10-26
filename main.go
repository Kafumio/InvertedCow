package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"InvertedCow/config"

	"github.com/gin-gonic/gin"
)

func newApp(engine *gin.Engine, config *config.AppConfig) *http.Server {
	srv := &http.Server{
		Addr:    config.Port,
		Handler: engine,
	}
	return srv
}

func main() {
	// 获取参数
	path, _ := os.Getwd()
	path = strings.ReplaceAll(path, "\\", "/")
	path = path + "/conf/config.ini"

	// 加载配置
	conf, err := config.InitSetting(path)
	if err != nil {
		log.Println("加载配置文件出错")
		return
	}

	if err != nil {
		log.Println(err)
	}

	// 注册路由
	srv, err := initApp(conf)
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Println(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown error: ", err)
	}
	log.Println("Server exiting")
}
