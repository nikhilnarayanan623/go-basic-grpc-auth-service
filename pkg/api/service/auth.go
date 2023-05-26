package service

import (
	"context"
	"net/http"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/api/service/interfaces"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/domain"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/pb/authpb"
	usecase "github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/usecase/interfaces"
)

type serviceServer struct {
	authpb.AuthServiceServer
	authUseCase usecase.AuthUseCase
}

func NewAuthServiceServer(useCase usecase.AuthUseCase) interfaces.ServiceServer {
	return &serviceServer{
		authUseCase: useCase,
	}
}

func (c *serviceServer) UserSignup(ctx context.Context, req *authpb.SignupRequest) (res *authpb.SignupResponse, err error) {

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

func (c *serviceServer) UserLogin(ctx context.Context, req *authpb.LoginRequest) (res *authpb.LoginResponse, err error) {
	return
}
func (c *serviceServer) ValidateAccessToken(ctx context.Context, req *authpb.ValidateRequest) (res *authpb.ValidateResponse, err error) {
	return
}
