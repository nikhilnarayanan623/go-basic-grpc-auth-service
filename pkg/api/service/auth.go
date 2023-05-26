package service

import (
	"context"
	"net/http"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/domain"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/pb/authpb"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/usecase/interfaces"
)

type ServiceServer struct {
	authpb.AuthServiceServer
	authUseCase interfaces.AuthUseCase
}

func NewAuthServiceServer(useCase interfaces.AuthUseCase) *ServiceServer {
	return &ServiceServer{
		authUseCase: useCase,
	}
}

func (c *ServiceServer) UserSignup(ctx context.Context, req *authpb.SignupRequest) (res *authpb.SignupResponse, err error) {

	_, err = c.authUseCase.UserSignup(context.Background(), domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})

	if err != nil {
		return &authpb.SignupResponse{
			Response: &authpb.Response{

				StatusCode: http.StatusBadRequest,
				Message:    "sign up failed",
				Error:      err.Error(),
			},
		}, nil
	}

	return &authpb.SignupResponse{
		Response: &authpb.Response{
			Message:    "successfully account created",
			StatusCode: http.StatusOK,
			Error:      "",
		},
	}, nil
}

func (c *ServiceServer) UserLogin(ctx context.Context, req *authpb.LoginRequest) (res *authpb.LoginResponse, err error) {
	return
}
func (c *ServiceServer) ValidateAccessToken(ctx context.Context, req *authpb.ValidateRequest) (res *authpb.ValidateResponse, err error) {
	return
}
