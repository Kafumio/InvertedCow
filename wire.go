//go:build wireinject
// +build wireinject

package main

import (
	"InvertedCow/config"
	"InvertedCow/controller"
	"InvertedCow/dao"
	"InvertedCow/data"
	"InvertedCow/interceptor"
	"InvertedCow/router"
	"InvertedCow/service"
	"github.com/google/wire"
	"net/http"
)

func initApp(*config.AppConfig) (*http.Server, error) {
	panic(wire.Build(
		dao.ProviderSet,
		data.ProviderSet,
		service.ProviderSet,
		controller.ProviderSet,
		interceptor.ProviderSet,
		router.SetupRouter,
		newApp),
	)
}
