package handler

import (
	"net/http"
	_ "os"

	"forge/internal/models"
	"forge/internal/service"

	"github.com/gin-gonic/gin"
)

type IAuthHandler interface {
	Login(c *gin.Context)
}

type AuthHandler struct {
	AuthService service.IAuthService
}

func NewAuthHandler(authService service.IAuthService) IAuthHandler {
	return &AuthHandler{authService}
}

// @Summary Login
// @Description Logs in a user
// @Tags auth
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} string "Logged in"
// @Failure 400 {string} string "Bad request"
// @Failure 401 {string} string "Unauthorized"
// @Router /login [post]
func (a AuthHandler) Login(c *gin.Context) {
	var loginData models.LoginRequest
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	tokenString, err := a.AuthService.Login(loginData)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
