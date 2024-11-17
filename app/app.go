package app

import (
	"inventory-app/internal/handler"
	"inventory-app/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetRouter(h *handler.AccHandler)*gin.Engine{
	r := gin.Default()
	r.POST("/api/register", h.RegisterHandler)
	r.POST("/api/login", h.LoginHandler)

	protected := r.Group("/account", middleware.Auth())
	{
		protected.GET(":username",h.GetAccount)
		protected.PUT("/",h.UpdateAccount)
		protected.DELETE("/",h.DeleteAccount)
	}
	return r
}