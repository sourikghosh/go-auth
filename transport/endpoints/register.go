package endpoints

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"unicode"

	"auth/implementation/auth"
	"auth/pkg"
	"auth/transport"

	"github.com/go-logr/logr"
)

// registerHandler handles the register route.
func registerHandler(s auth.Service, l logr.Logger) pkg.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var body transport.User
		var response transport.GenericResponse

		data, err := ioutil.ReadAll(request.(io.Reader))
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(data, &body)
		if err != nil {
			return response, pkg.AuthErr{
				Code: http.StatusBadRequest,
				Err:  err,
			}
		}

		if err := validateUser(ctx, body); err != nil {
			return nil, err
		}

		userID, err := s.Register(ctx, auth.User{
			UserName: body.UserName,
			FullName: body.FullName,
			Password: body.Password,
		})

		if err != nil {
			return nil, pkg.AuthErr{
				Code: http.StatusInternalServerError,
				Err:  err,
			}
		}

		response.Data = userID
		response.Message = "successfully created user"
		return response, nil
	}
}

// validateUser checks and validates all user attributes.
func validateUser(ctx context.Context, u transport.User) (err error) {
	if len([]rune(u.FullName)) <= 3 {
		err = errors.New("fullName cannot be less than 4 charecter")
	}

	if len([]rune(u.UserName)) <= 3 {
		err = errors.New("userName cannot be less than 4 charecter")
	}

	err = checkInvalidPassword(u.Password)

	if err == nil {
		return nil
	}

	return pkg.AuthErr{
		Code: http.StatusBadRequest,
		Err:  err,
	}
}

// checkInvalidPassword validates the incoming password with validation rules.
func checkInvalidPassword(password string) error {
	if len([]rune(password)) < 8 {
		return errors.New("password cannot be less than 8")
	}

	hasSymbols := false
	hasUpperL := false
	hasNumeric := false

	for _, char := range password {
		switch {
		case unicode.IsSymbol(char):
			hasSymbols = true

		case unicode.IsUpper(char):
			hasUpperL = true

		case unicode.IsDigit(char):
			hasNumeric = true
		}
	}

	if !hasNumeric || !hasUpperL || !hasSymbols {
		return errors.New("invalid passowrd should include atleast one speacial, upper and digit")
	}

	return nil
}
