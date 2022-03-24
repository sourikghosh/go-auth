package auth

import (
	"context"
	"errors"

	"auth/pkg"
	"auth/pkg/config"

	"github.com/go-logr/logr"
	"gorm.io/gorm"
)

type Service interface {
	Register(ctx context.Context, u User) (uint64, error)
	Login(ctx context.Context, u User) (string, error)
	Get(ctx context.Context, userID uint64) (User, error)
}

func NewService(logger logr.Logger, repo Repository, cfg config.AuthConfig, jsvc pkg.JWTService) Service {
	return &authSrv{
		l:    logger,
		r:    repo,
		cfg:  cfg,
		jsvc: jsvc,
	}
}

type authSrv struct {
	l    logr.Logger
	r    Repository
	cfg  config.AuthConfig
	jsvc pkg.JWTService
}

// Register service checks before creating a user if already exist
// and then computes a hash of the password and ceates a user.
func (s *authSrv) Register(ctx context.Context, u User) (uint64, error) {
	// checking if userExist with same userName
	usr, err := s.r.GetByFilter(map[string]interface{}{
		"user_name": u.UserName,
	})
	if !errors.Is(err, gorm.ErrRecordNotFound) || usr.ID != 0 {
		return 0, errors.New("user already exist")
	}

	// hashing the provided password for extra securtiy measures.
	hash, err := pkg.GenerateFromPassword(u.Password)
	if err != nil {
		return 0, err
	}

	// updating the raw password with encodedHash and creating the user
	u.Password = hash
	return s.r.Create(u)
}

// Login returns an access token if it fetches the user from the database and credentials match.
func (s *authSrv) Login(ctx context.Context, u User) (string, error) {
	// fetching the user from DB using userName
	usr, err := s.r.GetByFilter(map[string]interface{}{
		"user_name": u.UserName,
	})
	if err != nil {
		return "", err
	}

	// comparing the stored encodedHashed passowrd with provided password
	correct, err := pkg.ComparePasswordAndHash(u.Password, usr.Password)
	if err != nil || !correct {
		return "", err
	}

	// creates a JWT token using userID and sceret
	return s.jsvc.CreateJWTToken(usr.ID)
}

// Get fetches user by ID.
func (s *authSrv) Get(ctx context.Context, userID uint64) (User, error) {
	return s.r.GetByFilter(map[string]interface{}{
		"id": userID,
	})
}
