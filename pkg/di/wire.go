//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/api"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/api/service"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/clients"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/config"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/db"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/repository"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/token"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/usecase"
)

func InitiliazeService(cfg *config.Config) (*api.Server, error) {

	wire.Build(
		db.InitializeDatabase,
		token.NewJwtTokenAuth,
		repository.NewAuthRepository,
		clients.NewUserClient,
		usecase.NewAuthUseCase,
		service.NewAuthServiceServer,
		api.SetupAuthServer,
	)
	return &api.Server{}, nil
}
