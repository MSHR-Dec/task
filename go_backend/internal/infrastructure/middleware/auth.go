package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/MSHR-Dec/task/go_backend/internal/interface/controller"
	"github.com/MSHR-Dec/task/go_backend/pkg/oops"
)

func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if token := session.Get(controller.SessionKey); token == nil {
			oops.ResponseError(c, oops.Unauthorized{Message: "unauthorized user"})
			c.Abort()
		}
	}
}
