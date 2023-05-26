package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/domain"
	repo "github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/token"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/usecase/interfaces"
)

type authUseCase struct {
	authRepo  repo.AuthRepository
	tokenAuth token.TokenAuth
}

func NewAuthUseCase(repo repo.AuthRepository, tokenAuth token.TokenAuth) interfaces.AuthUseCase {
	return &authUseCase{
		authRepo:  repo,
		tokenAuth: tokenAuth,
	}
}
func (c *authUseCase) UserSignup(ctx context.Context, user domain.User) (userId uint32, err error) {

	return userId, nil
}

func (c *authUseCase) UserLogin(ctx context.Context, loginDetails domain.LoginRequest) (userId uint32, err error) {

	return userId, nil
}

func (c *authUseCase) GenerateAccessToken(ctx context.Context, userId uint32) (tokenString string, err error) {

	tokenString, err = c.tokenAuth.GenerateToken(token.TokenRequest{
		UserId:         userId,
		ExpirationTime: time.Now().Add(time.Minute * 30),
	})
	if err != nil {
		return tokenString, fmt.Errorf("failed to generate token \nerror:%s", err.Error())
	}
	return
}

func (c *authUseCase) VerifyAccessToken(ctx context.Context, tokenString string) (userId uint32, err error) {

	tokenRes, err := c.tokenAuth.VerifyToken(tokenString)
	if err != nil {
		return userId, fmt.Errorf("failed to verify token \nerror:%s", err.Error())
	}

	return tokenRes.UserId, nil
}
