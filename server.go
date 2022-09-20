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
	experienceRepository  repository.ExperienceRepository  = repository.NewExperienceRepository(db)
	skillrepository repository.SkillRepository = repository.NewSkillRepository(db)

	jwtService service.JWTService = service.NewJWTService()

	authService      service.AuthService      = service.NewAuthService(userRepository, applicantRepository, employeeRepository)
	applicantService service.ApplicantService = service.NewApplicantService(applicantRepository)
	employeeService service.EmployeeService = service.NewEmployeeService(employeeRepository)
	experienceService service.ExperienceService = service.NewExperienceService(experienceRepository)
	skillService service.SkillService = service.NewSkillService(skillrepository)

	authController      controller.AuthController      = controller.NewAuthController(authService, jwtService)
	applicantController controller.ApplicantController = controller.NewApplicantController(applicantService, jwtService)
	employeeController controller.EmployeeController = controller.NewEmployeeController(employeeService, jwtService)
	experienceController controller.ExperienceController = controller.NewExperienceController(experienceService, jwtService)
	skillController controller.SkillController = controller.NewSkillController(skillService, jwtService)
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

		authApplicantRoutes.POST("/jobexperiences", experienceController.CreateExperience)
		authApplicantRoutes.PUT("/jobexperiences/:id", experienceController.UpdateExperience)
		authApplicantRoutes.DELETE("/jobexperiences/:id", experienceController.DeleteExperience)
		authApplicantRoutes.GET("/jobexperiences/", experienceController.GetAllExperiences)
		authApplicantRoutes.GET("/jobexperiences/:id", experienceController.GetExperienceByID)
		
		authApplicantRoutes.POST("/skills", skillController.CreateSkill)
		authApplicantRoutes.PUT("/skills/:id", skillController.UpdateSkill)
		authApplicantRoutes.GET("/skills/:id", skillController.GetSkillByID)
		authApplicantRoutes.DELETE("/skills/:id", skillController.DeleteSkill)
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
