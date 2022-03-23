package auth

import "log"

type Service interface {
	Create(u User) (uint, error)
}

func NewService(logger log.Logger, repo Repository) Service {
	return &authSrv{
		l: logger,
		r: repo,
	}
}

type authSrv struct {
	l log.Logger
	r Repository
}

func (s *authSrv) Create(u User) (uint, error) {
	return s.r.Create(u)
}
