package middlewares

import (
	"context"
	"net/http"
	"paymentservice/internal/api/response"
	"paymentservice/internal/api_errors"

	"github.com/gin-gonic/gin"
)

func (e *GinMiddleware) Auth(authorization bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !authorization {
			c.Next()
			return
		}
		uid := c.Request.Header.Get("x-user-id")

		if uid == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.ResponseError{
				Message: "Unauthorized",
				Code:    api_errors.ErrUnauthorizedAccess,
			})
			return
		}

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "x-user-id", uid))

		c.Next()
	}
}
