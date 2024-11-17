package handler

import (
	"inventory-app/internal/model"
	"inventory-app/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccHandler struct {
	serv service.AccServiceMethod
}
func NewAccHandler(serv service.AccServiceMethod)*AccHandler{
	return &AccHandler{serv: serv}
}

func(h *AccHandler)RegisterHandler(c *gin.Context){
	var acc model.Account
	if err := c.ShouldBindJSON(&acc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	if err := h.serv.RegisterService(&acc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message":"account created"})
}

func(h *AccHandler)LoginHandler(c *gin.Context){
	var creds struct{
		Username string `json:"username"`
		Password string	`json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	token, err := h.serv.LoginService(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token":token})
}

func(h *AccHandler)DeleteAccount(c *gin.Context){
	var creds struct{
		Username string `json:"username"`
		Password string	`json:"password"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	if err := h.serv.DeleteAccountService(creds.Username,creds.Password); err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}

func (h *AccHandler)UpdateAccount(c *gin.Context) {
	var acc model.Account
	if err := c.ShouldBindJSON(&acc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedAcc, err := h.serv.UpdateAccountService(&acc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"account": updatedAcc})
}

func(h *AccHandler)GetAccount(c *gin.Context){
	username := c.Param("username")
	acc, err := h.serv.GetAccountByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"account": acc})
}