//+build wireinject

package main

import (
	"jeeo/api-boilerplate/configs"
	"jeeo/api-boilerplate/internal/application/service"
	"jeeo/api-boilerplate/internal/domain/repository"
	"jeeo/api-boilerplate/internal/entrypoint/handler"
	"jeeo/api-boilerplate/internal/entrypoint/http"
	"jeeo/api-boilerplate/internal/entrypoint/mapper"
	"jeeo/api-boilerplate/internal/infrastructure/persistence"

	"github.com/google/wire"
)

func initHttpService() http.HttpService {
	wire.Build(
		configs.ProvideConfig,
		persistence.ProvideDatabase,
		mapper.MapperSet,
		repository.RepositorySet,
		service.ServiceSet,
		handler.HandlerSet,
		http.HttpSet,
	)

	return nil
}
