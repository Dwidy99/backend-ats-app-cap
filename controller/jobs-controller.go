package controller

import (
	"errors"
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
	GetAllJobs(ctx *gin.Context)
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

func (c *jobsController) GetAllJobs(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		messError := fmt.Sprintf("Failed to get all jobs data, token user wrong or empty")
		response := helpers.BuildErrorResponse("Failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		messError := fmt.Sprintf("Failed to get all jobs data, token user wrong or empty")
		response := helpers.BuildErrorResponse("Failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if userID == 0 {
		errorMessage := gin.H{"error": errors.New("forbidden")}
		messError := fmt.Sprintf("Failed to get all jobs data, user with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("Failed to process request", messError, errorMessage)
		ctx.JSON(http.StatusForbidden, response)
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

	jobs, err := c.jobsService.AllJobs()
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}
	response := helpers.BuildResponse(true, "Success to get all jobs data", jobs)
	ctx.JSON(http.StatusOK, response)
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
