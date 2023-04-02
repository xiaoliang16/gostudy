package server

import (
	"github.com/gin-gonic/gin"
	"gostudy/api"
	"gostudy/middleware"
	"os"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Session(os.Getenv("SESSION_SERECT")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		v1.POST("user/register", api.UserRegister)

		v1.POST("user/login", api.UserLogin)

		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}

	return r
}