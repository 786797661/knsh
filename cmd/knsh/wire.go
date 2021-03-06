//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"knsh/internal/biz"
	"knsh/internal/conf"
	"knsh/internal/data"
	"knsh/internal/server"
	"knsh/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, *conf.Auth, *conf.Registry, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
