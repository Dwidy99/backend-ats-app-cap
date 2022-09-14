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
	db                  *gorm.DB                       = config.SetupConnectionDatabase()

	userRepository      repository.UserRepository      = repository.NewUserRepository(db)
	applicantRepository repository.ApplicantRepository = repository.NewApplicantRepository(db)
	employeeRepository  repository.EmployeeRepository  = repository.NewEmployeeRepository(db)

	jwtService service.JWTService = service.NewJWTService()

	authService      service.AuthService      = service.NewAuthService(userRepository, applicantRepository, employeeRepository)
	applicantService service.ApplicantService = service.NewApplicantService(applicantRepository)
	employeeService service.EmployeeService = service.NewEmployeeService(employeeRepository)

	authController      controller.AuthController      = controller.NewAuthController(authService, jwtService)
	applicantController controller.ApplicantController = controller.NewApplicantController(applicantService, jwtService)
	employeeController controller.EmployeeController = controller.NewEmployeeController(employeeService, jwtService)
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
	authApplicantRoutes := r.Group("/applicants", middleware.AuthorizeJWT(jwtService))
	{
		authApplicantRoutes.PUT("/users/:id", applicantController.EditApplicant)
		authApplicantRoutes.GET("/users/fetch", applicantController.FetchUserApplicant)
	}

	// Employees Routes
	authEmployeeRoutes := r.Group("/employees", middleware.AuthorizeJWT(jwtService))
	{
		authEmployeeRoutes.POST("/register", authController.RegisterEmployees)
		authEmployeeRoutes.PUT("/users/:id", employeeController.EditEmployee)
		authEmployeeRoutes.GET("/users/fetch", employeeController.FetchUserEmployee)
	}
	r.Run()
}
