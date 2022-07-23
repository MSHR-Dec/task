package adapter

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/MSHR-Dec/task/go_task/adapter/middleware"
)

func NewGin() *gin.Engine {
	r := gin.Default()

	r.Use(sessions.Sessions("task", NewRedisConnection()))
	setLogger(r)
	setCORS(r)
	setRoute(r)

	return r
}

func setLogger(r *gin.Engine) {
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, false))
	r.Use(ginzap.RecoveryWithZap(logger, true))
}

func setCORS(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: Environment.CorsAllowOrigins,
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Origin",
			"Access-Control-Allow-Headers",
			"Access-Control-Expose-Headers",
			"Access-Control-Allow-Credentials",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowCredentials: true,
	}))
}

func setRoute(r *gin.Engine) {
	taskController := injectTask()
	userController := injectUser()

	r.POST("/signup", userController.SignUp)
	r.POST("/signin", userController.SignIn)
	r.GET("/signout", userController.SignOut)

	profile := r.Group("/profile")
	{
		profile.Use(middleware.LoginRequired())
		profile.PUT("", userController.EditProfile)
	}

	user := r.Group("/users")
	{
		user.Use(middleware.LoginRequired())
		task := user.Group("/:userID/tasks")
		{
			task.POST("", taskController.Add)
			task.PUT("", taskController.Edit)
			task.GET("", taskController.ListByUserID)
		}
	}
}
