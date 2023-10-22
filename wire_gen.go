// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"InvertedCow/config"
	"InvertedCow/controller"
	"InvertedCow/dao"
	"InvertedCow/data"
	"InvertedCow/interceptor"
	"InvertedCow/router"
	"InvertedCow/service"
	"net/http"
)

// Injectors from wire.go:

func initApp(appConfig *config.AppConfig) (*http.Server, error) {
	db := data.NewGormClient(appConfig)
	userDao := dao.NewUserDao()
	accountService := service.NewAccountService(db, userDao)
	accountController := controller.NewAccountController(accountService)
	controllerController := controller.NewController(accountController)
	corsInterceptor := interceptor.NewCorsInterceptor()
	engine := router.SetupRouter(controllerController, corsInterceptor)
	server := newApp(engine, appConfig)
	return server, nil
}