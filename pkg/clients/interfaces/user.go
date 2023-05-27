package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/domain"
)

type UserClient interface {
	FindUserByEmail(ctx context.Context, email string) (domain.User, error)
	SaveUser(ctx context.Context, user domain.User) (userId uint32, err error)
}
