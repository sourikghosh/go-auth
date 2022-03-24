package http

import (
	"net/http"
	"strings"

	"auth/pkg"
	"auth/transport/endpoints"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

// NewHTTPService takes all the endpoints and returns handler.
func NewHTTPService(endpoints endpoints.Endpoints, svc pkg.JWTService) http.Handler {

	r := gin.New()

	r.HandleMethodNotAllowed = true
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("/auth-svc/v1")
	{
		auth.POST("/register", endpointRequestEncoder(endpoints.Register))
		auth.POST("/login", endpointRequestEncoder(endpoints.Login))

		auth.Use(authZMiddleware(svc))
		auth.GET("/profile", endpointRequestEncoder(endpoints.GetProfile))
	}

	return r
}

// endpointRequestEncoder encodes request and does error handling
// and send response.
func endpointRequestEncoder(endpoint pkg.Endpoint) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var statusCode int
		// process the request with its handler
		response, err := endpoint(c, c.Request.Body)

		if err != nil {
			// if statusCode is not send then return InternalServerErr
			switch e := err.(type) {
			case pkg.Error:
				statusCode = e.Status()

			default:
				statusCode = http.StatusInternalServerError
			}

			c.AbortWithStatusJSON(statusCode, gin.H{
				"error":   true,
				"message": err.Error(),
			})

			return
		}

		// if err did not occur then return Ok status
		c.JSON(http.StatusOK, response)
	}

	return gin.HandlerFunc(fn)
}

// authZMiddleware checks if a user is part of an organization and authenticates the token.
func authZMiddleware(svc pkg.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Missing token",
			})

			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := svc.Verify(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid JWT token",
			})

			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid JWT token",
			})

			return
		}

		userID, exist := claims["user_id"].(float64)
		if !exist {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid JWT token",
			})

			return
		}

		// Set the current user in context.
		c.Set("user_id", uint64(userID))
		c.Set("token", tokenString)
	}
}
