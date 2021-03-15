//+build wireinject

package main

import (
	"github.com/Jeeo/golang-ddd-boilerplate/configs"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/application"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/application/repository"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/handler"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/http"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/entrypoint/mapper"
	"github.com/Jeeo/golang-ddd-boilerplate/internal/infrastructure/persistence"

	"github.com/google/wire"
)

func initHttpService() http.HttpService {
	wire.Build(
		configs.ProvideConfig,
		persistence.ProvideDatabase,
		mapper.MapperSet,
		repository.RepositorySet,
		application.ApplicationSet,
		handler.HandlerSet,
		http.HttpSet,
	)

	return nil
}
