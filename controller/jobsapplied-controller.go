package controller

import (
	"errors"
	"fmt"
	"mini-project/helpers"
	"mini-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type JobAppliedController interface {
	JobsAppliedByApplicantID(ctx *gin.Context)
}

type jobAppliedController struct {
	jobAppliedService service.JobAppliedService
	jwtService service.JWTService
}

func NewJobAppliedController(jobApplied service.JobAppliedService, jwtService service.JWTService) JobAppliedController {
	return &jobAppliedController{
		jobAppliedService: jobApplied,
		jwtService: jwtService,
	}
}

func (c *jobAppliedController) JobsAppliedByApplicantID(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), err)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if userID == 0 {
		errorMessage := gin.H{"error": errors.New("forbidden")}
		messError := fmt.Sprintf("user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, errorMessage)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	applicant, err := c.jobAppliedService.GetApplicantByID(userID)

	applied, err := c.jobAppliedService.JobAppliedByApplicantID(int(applicant.ID))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	response := helpers.BuildResponse(true, "success to get job applied", applied)
	ctx.JSON(http.StatusOK, response)
}