package handler

import (
	"net/http"
	"os"

	"forge/internal/models"
	"forge/internal/service"

	"github.com/gin-gonic/gin"
)

type IAuthHandler interface {
	Login(c *gin.Context)
	Profile(c *gin.Context)
}

type AuthHandler struct {
	AuthService service.IAuthService
}

func NewAuthHandler(authService service.IAuthService) IAuthHandler {
	return &AuthHandler{authService}
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param loginData body models.LoginRequest true "Login data"
// @Success 200 {object} models.TokenResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
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

	domain := os.Getenv("COOKIE_DOMAIN")

	c.SetCookie("jwt", tokenString, 3600, "/", domain, true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Successful"})
}

// Profile handles the request to retrieve the profile information of the authenticated user.
// @Summary Retrieve user profile
// @Description Get the profile information of the authenticated user
// @Tags profile
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Authorized"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Router /profile [get]
func (a AuthHandler) Profile(c *gin.Context) {
	// Extraer el username desde el contexto
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Devolver los datos del usuario (esto puede venir de una BD en un caso real)
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"message":  "Authorized",
	})
}
