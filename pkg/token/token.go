package token

import "time"

type TokenAuth interface {
	GenerateToken(tokenReq TokenRequest) (tokenString string, err error)
	VerifyToken(tokenString string) (TokenResponse, error)
}

type TokenRequest struct {
	UserId         uint32
	ExpirationTime time.Time
}

type TokenResponse struct {
	UserId uint32
}
