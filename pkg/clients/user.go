package clients

import (
	"context"
	"fmt"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/clients/interfaces"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/config"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/domain"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type userClient struct {
	client pb.UserServiceClient
}

func NewUserClient(cfg *config.Config) (interfaces.UserClient, error) {

	gClient, err := grpc.Dial(cfg.UserClientPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect with user service")
	}
	client := pb.NewUserServiceClient(gClient)

	return &userClient{
		client: client,
	}, nil
}

func (c *userClient) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {

	res, err := c.client.FindUserByEmail(ctx, &pb.FindUserByEmailRequest{Email: email})
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:        res.UserId,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
		Password:  res.Password,
	}, nil
}
func (c *userClient) SaveUser(ctx context.Context, user domain.User) (userId uint32, err error) {

	res, err := c.client.SaveUser(ctx, &pb.SaveUserRequest{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	})
	if err != nil {
		return 0, err
	}

	return res.UserId, nil
}
