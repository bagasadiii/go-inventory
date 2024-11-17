package handler

import "github.com/gin-gonic/gin"

type AccHandlerMethod interface {
	RegisterHandler(c *gin.Context)
	LoginHandler(c *gin.Context)
	UpdateAccount(c *gin.Context)
	DeleteAccount(c *gin.Context)
	GetAccount(c *gin.Context)
}