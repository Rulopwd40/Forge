package handler

import (
	"net/http"

	"forge/internal/models"
	"forge/internal/service"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	RegisterUser(c *gin.Context)
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
// @Success 200 {object} models.User
// @Router /user/register [post]
func (h *UserHandler) RegisterUser(c *gin.Context) {
	var userStruct models.User

	if err := c.ShouldBindJSON(&userStruct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind JSON",
		})
		return
	}

	err := h.UserService.RegisterUser(userStruct)

	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "User/email already exists",
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Bad request",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user":    userStruct,
	})
}
