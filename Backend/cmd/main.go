// @title Forge API
// @version 1.0
// @description API for user management in Forge
// @host localhost:8080
// @BasePath /
package main

import (
	_ "forge/docs"
	_ "forge/internal/handler" // Asegúrate de que esté importado
	_ "forge/internal/repository"
	"forge/internal/server"
	_ "forge/internal/service" // También es importante que los servicios estén importados
)

func main() {
	server.LoadEnvVariables()
	server.Run()
}
