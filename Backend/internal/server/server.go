// @host localhost:8080
// @BasePath /
package server

import (
	"forge/internal/database"
	"forge/internal/handler"
	"forge/internal/repository"
	"forge/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func Run() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(CORSMiddleware())

	server := &Server{
		db: database.ConnectDB(),
	}

	userHandler := handler.NewUserHandler(service.NewUserService(repository.NewUserRepository(server.db)))
	SetupUserRoutes(r, userHandler)
	// Inicia el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupUserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {

	print(userHandler)
	r.POST("/user/register", userHandler.RegisterUser)
	r.GET("/user", handler.GetUsers)
}
