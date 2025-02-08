package handler

import (
	"net/http"
	"os"
	_ "os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary Get Swagger JSON
// @Description Get Swagger JSON
// @Tags Swagger
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /swagger/index.json [get]
func SwaggerJSON(c *gin.Context) {
	c.JSON(http.StatusOK, ginSwagger.URL(os.Getenv("SWAGGER_URL")))
}
