package infrastructure

import (
	"github.com/MSHR-Dec/task/go_backend/internal/infrastructure/middleware"
	"time"

	"github.com/gin-contrib/sessions"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewGin() *gin.Engine {
	r := gin.Default()

	r.Use(sessions.Sessions("task", NewRedisConnection()))
	setLogger(r)
	setRoute(r)

	return r
}

func setLogger(r *gin.Engine) {
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, false))
	r.Use(ginzap.RecoveryWithZap(logger, true))
}

func setRoute(r *gin.Engine) {
	hello := r.Group("/")
	{
		hello.Use(middleware.LoginRequired())
		hello.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello!"})
		})
	}

	user := injectUser()

	r.POST("/signin", user.SignIn)
	r.POST("/signup", user.SignUp)
	r.GET("/signout", user.SignOut)
}
