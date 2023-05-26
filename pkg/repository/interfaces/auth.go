package interfaces

import (
	"context"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/domain"
)

type AuthRepository interface {
	FindUserByEmail(ctx context.Context, email string) (user domain.User, err error)
	SaveUser(ctx context.Context, user domain.User) (userId uint32, err error)
}
