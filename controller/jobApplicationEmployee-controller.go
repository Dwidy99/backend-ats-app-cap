package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PutraFajarF/backend-ats-app-cap/dto"
	"github.com/PutraFajarF/backend-ats-app-cap/helpers"
	"github.com/PutraFajarF/backend-ats-app-cap/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JobApplicationEmployeeController interface {
	ProgressApplication(ctx *gin.Context)
}

type jobApplicationEmployeeController struct {
	serviceJobApplicationEmployee service.JobApplicationEmployeeService
	JWTService                    service.JWTService
}

func NewJobApplicationEmployeeController(serviceJobApplicationEmployee service.JobApplicationEmployeeService, jwtService service.JWTService) JobApplicationEmployeeController {
	return &jobApplicationEmployeeController{
		serviceJobApplicationEmployee: serviceJobApplicationEmployee,
		JWTService:                    jwtService,
	}
}

func (c *jobApplicationEmployeeController) ProgressApplication(ctx *gin.Context) {
	var inputID dto.GetJobApplicationEmployee

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed update progress application", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData dto.UpdateJobApplicationEmployeeDTO
	err = ctx.ShouldBind(&inputData)
	if err != nil {
		response := helpers.BuildErrorResponse("failed update progress application", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.JWTService.ValidateToken(authHeader)
	if errToken != nil {
		response := helpers.BuildErrorResponse("failed to process request", errToken.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, _ := c.serviceJobApplicationEmployee.GetUserByID(userID)
	if user.Role != "admin" {
		response := helpers.BuildErrorResponse("failed to process request", "role is not admin", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	jobApplicantion, err := c.serviceJobApplicationEmployee.UpdateProgress(inputData, int(inputID.ID))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "success to apply new job", jobApplicantion)
	ctx.JSON(http.StatusOK, response)
}
