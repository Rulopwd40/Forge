package handler

import (
	"net/http"

	"forge/internal/models"
	"forge/internal/service"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	RegisterUser(c *gin.Context)
	GetUserData(c *gin.Context)
}

type UserHandler struct {
	UserService service.IUserService
}

func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// RegisterUser godoc
// @Summary Registers a new user
// @Description This endpoint allows to register a new user.
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body models.User true "User data"
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 409 {object} models.ErrorResponse
// @Router /user/register [post]
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var userStruct models.User

	if err := c.ShouldBindJSON(&userStruct); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Failed to bind JSON"})
		return
	}

	err := h.UserService.RegisterUser(userStruct)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusConflict, models.ErrorResponse{Error: "User/email already exists"})
		} else {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Bad request"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user":    userStruct,
	})
}

// GetUserData godoc
// @Summary Get user data
// @Description Retrieves user data based on the provided username
// @Tags user
// @Accept json
// @Produce json
// @Param username query string true "Username"
// @Success 200 {object} models.UserResponse "User data retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Username is required"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Router /user [get]
func (h *UserHandler) GetUserData(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Username is required"})
		return
	}

	user, err := h.UserService.GetUserData(username)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{Error: "User not found"})
		return
	}

	userResponse := models.UserResponse{
		Username: user.Username,
		Email:    user.Email,
		Level:    user.Level,
		Name:     user.Name,
		Message:  "User retrieved successfully",
	}
	c.JSON(http.StatusOK, userResponse)
}
