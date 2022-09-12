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
	applicantRepository repository.ApplicantRepository = repository.NewApplicantRepository(db)
	employeeRepository repository.EmployeeRepository = repository.NewEmployeeRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository, applicantRepository, employeeRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {

	defer config.CloseConnectionDatabase(db)
	r := gin.Default()

	authRoutes := r.Group("/")
	{
		authRoutes.POST("login", authController.Login)
		authRoutes.POST("applicants/register", authController.RegisterApplicants)
	}

	// Applicant Routes

	// Employees Routes
	authEmployeeRoutes := r.Group("/employees", middleware.AuthorizeJWT(jwtService))
	{
		authEmployeeRoutes.POST("/register", authController.RegisterEmployees)
	}

	r.Run()
}
