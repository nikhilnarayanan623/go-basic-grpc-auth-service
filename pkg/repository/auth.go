package repository

import (
	"context"

	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/domain"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type authDatabase struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.AuthRepository {
	return &authDatabase{
		db: db,
	}
}

func (c *authDatabase) FindUserByEmail(ctx context.Context, email string) (user domain.User, err error) {
	query := `SELECT * FROM users WHERE email = $1`
	err = c.db.Raw(query, email).Scan(&user).Error
	return
}
func (c *authDatabase) SaveUser(ctx context.Context, user domain.User) (userId uint32, err error) {
	query := `INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4 ) RETURNING id AS user_id`
	err = c.db.Raw(query, user.FirstName, user.LastName, user.Email, user.Password).Scan(&userId).Error
	return
}
