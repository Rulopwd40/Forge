package handler

import (
	"net/http"

	"forge/internal/models"
	"forge/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
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
// @Success 200 {object} models.User
// @Router /user/register [post]
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var userStruct models.User

	if err := c.ShouldBindJSON(&userStruct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.UserService.RegisterUser(userStruct)

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user":    userStruct,
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get users called",
	})
}
