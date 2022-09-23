package controller

import (
	"fmt"
	"mini-project/dto"
	"mini-project/helpers"
	"mini-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JobApplicationController interface {
	CreateApply(ctx *gin.Context)
}

type jobApplicationController struct {
	serviceJobApplication service.JobApplicationService
	JWTService service.JWTService
}

func NewJobApplicationController(serviceJobApplication service.JobApplicationService, jwtService service.JWTService) JobApplicationController {
	return &jobApplicationController{
		serviceJobApplication: serviceJobApplication,
		JWTService: jwtService,
	}
}

func (c *jobApplicationController) CreateApply(ctx *gin.Context) {
	var inputData dto.CreateJobApplication

	err := ctx.ShouldBind(&inputData)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
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

	user, err := c.serviceJobApplication.GetUserByID(userID)
	if user.Role != "user" {
		response := helpers.BuildErrorResponse("failed to process request", "role is not user", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	applicant, err := c.serviceJobApplication.GetApplicantByID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	fmt.Println(inputData)

	jobApplicantion, err := c.serviceJobApplication.CreateJobApplicant(inputData, int(applicant.ID))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "success to create job skill", jobApplicantion)
	ctx.JSON(http.StatusOK, response)
}