package endpoints

import (
	"auth/implementation/auth"
	"auth/pkg"

	"github.com/go-logr/logr"
)

// Endpoints exposes all endpoints.
type Endpoints struct {
	Register   pkg.Endpoint
	Login      pkg.Endpoint
	GetProfile pkg.Endpoint
}

// MakeEndpoints initialises all required services and composes all endpoints.
func MakeEndpoints(svc auth.Service, l logr.Logger) Endpoints {
	return Endpoints{
		Register:   registerHandler(svc, l),
		Login:      loginHandler(svc, l),
		GetProfile: profileHandler(svc, l),
	}
}
