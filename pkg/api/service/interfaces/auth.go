package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/pb/authpb"
)

type ServiceServer interface {
	UserSignup(ctx context.Context, req *authpb.SignupRequest) (res *authpb.SignupResponse, err error)
	UserLogin(ctx context.Context, req *authpb.LoginRequest) (res *authpb.LoginResponse, err error)
	ValidateAccessToken(ctx context.Context, req *authpb.ValidateRequest) (res *authpb.ValidateResponse, err error)
	authpb.AuthServiceServer
}
