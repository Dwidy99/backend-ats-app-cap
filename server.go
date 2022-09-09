package main

import (
	"mini-project/config"
	"mini-project/controller"
	"mini-project/middleware"
	"mini-project/repository"
	"mini-project/service"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	db             *gorm.DB                  = config.SetupConnectionDatabase()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {

	defer config.CloseConnectionDatabase(db)
	r := gin.Default()

	authRoutes := r.Group("/")
	{
		authRoutes.POST("login", authController.Login)
	}

	// Applicants
	authApplicantsRoutes := r.Group("/applicants")
	{
		authApplicantsRoutes.POST("/register", authController.RegisterApplicants)
	}

	// Employees
	authEmployeeRoutes := r.Group("/employees", middleware.AuthorizeJWT(jwtService))
	{
		authEmployeeRoutes.POST("/register", authController.RegisterEmployee)
	}

	r.Run()
}
