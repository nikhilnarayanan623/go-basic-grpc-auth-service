package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nikhilnarayanan623/go-basic-grpc-auth-service/pkg/config"
)

type jwtAuth struct {
	jwtSecret string
}

type JwtClaims struct {
	jwt.RegisteredClaims
	UserId uint32
}

func NewJwtTokenAuth(cfg *config.Config) TokenAuth {
	return &jwtAuth{
		jwtSecret: cfg.JWTSecret,
	}
}

func (c *jwtAuth) GenerateToken(tokenReq TokenRequest) (tokenString string, err error) {

	claims := &JwtClaims{
		UserId: tokenReq.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenReq.ExpirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(c.jwtSecret))
	return
}
func (c *jwtAuth) VerifyToken(tokenString string) (TokenResponse, error) {

	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method on token")
		}
		return []byte(c.jwtSecret), nil
	})

	if err != nil {
		return TokenResponse{}, fmt.Errorf("faild to parse token string to token")
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return TokenResponse{}, fmt.Errorf("faild to parse the token")
	}

	return TokenResponse{
		UserId: claims.UserId,
	}, nil
}
