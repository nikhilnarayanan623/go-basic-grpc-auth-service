package usecase

import (
	"context"
	"fmt"
	"time"

	client "github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/clients/interfaces"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/domain"
	repo "github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/token"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/utils"
)

type authUseCase struct {
	authRepo   repo.AuthRepository
	tokenAuth  token.TokenAuth
	userClient client.UserClient
}

func NewAuthUseCase(repo repo.AuthRepository, tokenAuth token.TokenAuth, userClient client.UserClient) interfaces.AuthUseCase {
	return &authUseCase{
		authRepo:   repo,
		tokenAuth:  tokenAuth,
		userClient: userClient,
	}
}
func (c *authUseCase) UserSignup(ctx context.Context, user domain.User) (userId uint32, err error) {

	existUser, err := c.userClient.FindUserByEmail(context.Background(), user.Email)
	if err != nil {
		return
	} else if existUser.ID != 0 {
		return userId, fmt.Errorf("user already exist with given email")
	}

	user.Password, err = utils.GenerateHashFromPassword(user.Password)
	if err != nil {
		return 0, fmt.Errorf("faild to hash user password error:%s", err.Error())
	}

	userId, err = c.userClient.SaveUser(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("failed to save user details \nerror:%s", err.Error())
	}

	return userId, nil
}

func (c *authUseCase) UserLogin(ctx context.Context, loginDetails domain.LoginRequest) (userId uint32, err error) {

	user, err := c.userClient.FindUserByEmail(ctx, loginDetails.Email)
	if err != nil {
		return 0, fmt.Errorf("failed to find user by email %s", err.Error())
	} else if user.ID == 0 {
		return 0, fmt.Errorf("user not found with given details")
	}

	if !utils.VerifyHashAndPassword(user.Password, loginDetails.Password) {
		return 0, fmt.Errorf("user password doesn't match")
	}

	return user.ID, nil
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
