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

type JobsController interface {
	CreatedJobs(ctx *gin.Context)
}

type jobsController struct {
	jobsService service.JobsService
	jwtService  service.JWTService
}

func NewJobsController(jobServ service.JobsService, jwtService service.JWTService) JobsController {
	return &jobsController{
		jobsService: jobServ,
		jwtService:  jwtService,
	}
}

func (c *jobsController) CreatedJobs(ctx *gin.Context) {
	var jobsInput dto.CreateJobsDTO
	err := ctx.ShouldBind(&jobsInput)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		messError := fmt.Sprintf("failed to create new job, token user admin wrong or empty")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		messError := fmt.Sprintf("failed to create new job, user admin with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.jobsService.GetUserByID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if user.Role != "admin" {
		messError := fmt.Sprintf("failed to create jobs, role is not admin")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	jobs, err := c.jobsService.CreateJobs(jobsInput, userID)
	if err != nil {
		messError := fmt.Sprintf("failed to create job, user admin with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "success to create job", jobs)
	ctx.JSON(http.StatusOK, response)
}
