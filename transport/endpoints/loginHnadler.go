package endpoints

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"auth/implementation/auth"
	"auth/pkg"
	"auth/transport"

	"github.com/gin-gonic/gin"
	"github.com/go-logr/logr"
)

func loginHandler(s auth.Service, l logr.Logger) pkg.Endpoint {
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

		token, err := s.Login(ctx, auth.User{
			UserName: body.UserName,
			Password: body.Password,
		})

		if err != nil {
			l.Error(err, "login attempt failed")

			return nil, pkg.AuthErr{
				Code: http.StatusForbidden,
				Err:  errors.New("invalid username/password"),
			}
		}

		response.Data = token
		response.Message = "OK"
		return response, nil
	}
}

func profileHandler(s auth.Service, l logr.Logger) pkg.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		fmt.Println("here reached")
		var response transport.GenericResponse

		c, ok := ctx.(*gin.Context)
		if !ok {
			return nil, pkg.AuthErr{Code: http.StatusInternalServerError,
				Err: errors.New("error typecasting gin context")}
		}

		v, exist := c.Get("user_id")
		if !exist {
			return nil, pkg.AuthErr{Code: http.StatusForbidden,
				Err: errors.New("userID is missing")}
		}

		userID, ok := v.(uint64)
		if !ok || userID == 0 {
			return nil, pkg.AuthErr{Code: http.StatusForbidden,
				Err: errors.New("invalid userID found")}
		}

		usr, err := s.Get(ctx, userID)
		if err != nil {
			return nil, pkg.AuthErr{
				Code: http.StatusInternalServerError,
				Err:  err,
			}
		}

		response.Data = transport.User{
			ID:        userID,
			CreatedAt: usr.CreatedAt,
			UpdatedAt: usr.UpdatedAt,
			UserName:  usr.UserName,
			FullName:  usr.FullName,
		}

		response.Message = "Ok"

		return response, nil
	}
}
