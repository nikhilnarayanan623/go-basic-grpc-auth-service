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

	userID, err := c.authUseCase.UserLogin(ctx, domain.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return &pb.LoginResponse{
			Response: &pb.Response{
				StatusCode: http.StatusBadRequest,
				Message:    "failed to login",
				Error:      err.Error(),
			},
		}, nil
	}

	accessToken, err := c.authUseCase.GenerateAccessToken(ctx, userID)
	if err != nil {
		return &pb.LoginResponse{
			Response: &pb.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "failed to generate access token",
				Error:      err.Error(),
			},
		}, nil
	}

	return &pb.LoginResponse{
		Response: &pb.Response{
			StatusCode: http.StatusOK,
			Message:    "successfully user validated and access token generated",
			Error:      "",
		},
		AccessToken: accessToken,
	}, nil
}
func (c *serviceServer) ValidateAccessToken(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {

	userID, err := c.authUseCase.VerifyAccessToken(context.Background(), req.AccessToken)
	if err != nil {
		return &pb.ValidateResponse{
			Response: &pb.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "failed to verify token",
				Error:      err.Error(),
			},
		}, nil
	}
	return &pb.ValidateResponse{
		Response: &pb.Response{
			StatusCode: http.StatusOK,
			Message:    "successfully token verified",
			Error:      "",
		},
		UserId: userID,
	}, nil
}
