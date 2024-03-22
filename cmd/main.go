package main

import (
	"project-structure/internal/handler"
	"project-structure/internal/infrastructure"
	"project-structure/internal/repository"
	"project-structure/internal/router"
	"project-structure/internal/service"

	_ "project-structure/cmd/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			GO PROJECT STRUCTURE
// @version		2.0
// @description	golang project structure
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
// @schemes		http
func main() {
	g := gin.Default()

	usersGroup := g.Group("/users")

	// dependency injection
	gorm := infrastructure.NewGormPostgres()
	userRepo := repository.NewUserQuery(gorm)
	userSvc := service.NewUserService(userRepo)
	userHdl := handler.NewUserHandler(userSvc)
	userRouter := router.NewUserRouter(usersGroup, userHdl)

	
	// mount
	userRouter.Mount()
	// swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run(":3000")
}