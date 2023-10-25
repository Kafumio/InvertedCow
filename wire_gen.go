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
	"InvertedCow/router"
	"InvertedCow/service"
	"net/http"
)

// Injectors from wire.go:

func initApp(appConfig *config.AppConfig) (*http.Server, error) {
	db := data.NewGormClient(appConfig)
	client := data.NewRedisClient(appConfig)
	userDao := dao.NewUserDao()
	accountService := service.NewAccountService(appConfig, db, client, userDao)
	accountController := controller.NewAccountController(accountService)
	relationService := service.NewRelationService(db, client, userDao)
	relationController := controller.NewRelationController(relationService)
	controllerController := controller.NewController(accountController, relationController)
	engine := router.SetupRouter(controllerController)
	server := newApp(engine, appConfig)
	return server, nil
}
