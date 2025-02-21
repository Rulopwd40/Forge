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
	Logout(c *gin.Context)
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
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request"})
		return
	}
	tokenString, err := a.AuthService.Login(loginData)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Unauthorized"})
		return
	}

	domain := os.Getenv("COOKIE_DOMAIN")
	if domain == "" {
		domain = "localhost"
	}

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
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"message":  "Authorized",
	})
}

// Logout godoc
// @Summary Logout user
// @Description Clear the JWT cookie
// @Tags auth
// @Success 200 {object} map[string]string "message"
// @Router /logout [post]
func (a AuthHandler) Logout(c *gin.Context) {
	domain := os.Getenv("COOKIE_DOMAIN")
	if domain == "" {
		domain = "localhost"
	}

	// Eliminar la cookie estableciendo su valor vacío y expirando inmediatamente
	c.SetCookie("jwt", "", -1, "/", domain, true, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
