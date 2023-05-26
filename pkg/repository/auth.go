package repository

import (
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
