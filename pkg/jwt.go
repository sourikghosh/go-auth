package pkg

import (
	"auth/pkg/config"
	"errors"
	"time"

	"github.com/go-logr/logr"
	jwt "github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	CreateJWTToken(userid uint64) (string, error)
	Verify(tokenString string) (*jwt.Token, error)
}

func NewJWTService(cfg config.AuthConfig, l logr.Logger) JWTService {
	return &jwtAuth{
		cfg: cfg,
		l:   l,
	}
}

type jwtAuth struct {
	cfg config.AuthConfig
	l   logr.Logger
}

//CreateJWTToken creates new JWT Access Token
func (j *jwtAuth) CreateJWTToken(userid uint64) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(j.cfg.JWTSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// Verify take token as string.
func (j *jwtAuth) Verify(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signed jwt")
		}

		return []byte(j.cfg.JWTSecret), nil
	})
}
