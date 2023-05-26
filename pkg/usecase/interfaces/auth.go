package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/domain"
)

type AuthUseCase interface {
	UserLogin(ctx context.Context, user domain.LoginRequest) (userId uint32, err error)
	GenerateAccessToken(ctx context.Context, userId uint32) (tokenString string, err error)
	VerifyAccessToken(ctx context.Context, tokenString string) (userId uint32, err error)
	UserSignup(ctx context.Context, user domain.User) (userId uint32, err error)
}
