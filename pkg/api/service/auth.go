package service

import (
	"context"
	"net/http"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/domain"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/pb"
	usecase "github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/usecase/interfaces"
)

type serviceServer struct {
	pb.AuthServiceServer
	authUseCase usecase.AuthUseCase
}

func NewAuthServiceServer(useCase usecase.AuthUseCase) pb.AuthServiceServer {
	return &serviceServer{
		authUseCase: useCase,
	}
}

func (c *serviceServer) UserSignup(ctx context.Context, req *pb.SignupRequest) (res *pb.SignupResponse, err error) {

	_, err = c.authUseCase.UserSignup(context.Background(), domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	})

	if err != nil {
		return &pb.SignupResponse{
			Response: &pb.Response{

				StatusCode: http.StatusBadRequest,
				Message:    "sign up failed",
				Error:      err.Error(),
			},
		}, nil
	}

	return &pb.SignupResponse{
		Response: &pb.Response{
			Message:    "successfully account created",
			StatusCode: http.StatusOK,
			Error:      "",
		},
	}, nil
}

func (c *serviceServer) UserLogin(ctx context.Context, req *pb.LoginRequest) (res *pb.LoginResponse, err error) {
	return
}
func (c *serviceServer) ValidateAccessToken(ctx context.Context, req *pb.ValidateRequest) (res *pb.ValidateResponse, err error) {
	return
}
