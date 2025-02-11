// @host localhost:8080
// @BasePath /
package server

import (
	"forge/internal/database"
	"forge/internal/handler"
	"forge/internal/repository"
	"forge/internal/service"

	"os"

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

	Initialize(r, server)

	// Inicia el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowOrigin := os.Getenv("ALLOW_ORIGIN")
		if allowOrigin == "" {
			allowOrigin = "*"
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
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

func Initialize(r *gin.Engine, server *Server) {

	//Repository
	userRepository := repository.NewUserRepository(server.db)

	//Service
	passwordService := service.NewPasswordService()
	userService := service.NewUserService(userRepository, passwordService)
	authService := service.NewAuthService(userService, passwordService)

	//Handler
	userHandler := handler.NewUserHandler(userService)
	SetupUserRoutes(r, userHandler)

	authHandler := handler.NewAuthHandler(authService)
	SetupAuthRoutes(r, authHandler)
}

func SetupUserRoutes(r *gin.Engine, userHandler handler.IUserHandler) {
	r.POST("/user/register", userHandler.RegisterUser)
}

func SetupAuthRoutes(r *gin.Engine, authHandler handler.IAuthHandler) {
	r.POST("/login", authHandler.Login)
}
