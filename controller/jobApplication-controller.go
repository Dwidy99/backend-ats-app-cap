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

type JobApplicationController interface {
	CreateApply(ctx *gin.Context)
}

type jobApplicationController struct {
	serviceJobApplication service.JobApplicationService
	JWTService            service.JWTService
}

func NewJobApplicationController(serviceJobApplication service.JobApplicationService, jwtService service.JWTService) JobApplicationController {
	return &jobApplicationController{
		serviceJobApplication: serviceJobApplication,
		JWTService:            jwtService,
	}
}

func (c *jobApplicationController) CreateApply(ctx *gin.Context) {
	var inputData dto.CreateJobApplicationDTO

	err := ctx.ShouldBind(&inputData)
	if err != nil {
		res := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
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

	role := fmt.Sprintf("%v", claims["role"])
	if role != "user" || role == "" {
		errMessage := fmt.Sprintf("role is not %v", role)
		response := helpers.BuildErrorResponse("failed to process request", errMessage, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	applicant, err := c.serviceJobApplication.GetApplicantByID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	jobApplicantion, err := c.serviceJobApplication.CreateJobApplicant(inputData, int(applicant.ID))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "success to apply new job", jobApplicantion)
	ctx.JSON(http.StatusOK, response)
}
