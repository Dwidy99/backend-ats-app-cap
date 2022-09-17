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

type ExperienceController interface {
	CreateExperience(ctx *gin.Context)
	UpdateExperience(ctx *gin.Context)
}

type experienceController struct {
	experienceService service.ExperienceService
	jwtService service.JWTService
}

func NewExperienceController(expService service.ExperienceService, jwtService service.JWTService) ExperienceController {
	return &experienceController{
		experienceService: expService,
		jwtService: jwtService,
	}
}

func (c *experienceController) CreateExperience(ctx *gin.Context) {
	var experienceInput dto.CreateExperienceDTO

	err := ctx.ShouldBind(&experienceInput)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		messError := fmt.Sprintf("failed to access update job experience, token user applicant wrong or empty")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		messError := fmt.Sprintf("failed to access create job experience, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	user, err := c.experienceService.GetUserByID(userID)
	if user.Role != "user" {
		messError := fmt.Sprintf("failed to access create job experience, role is unknow")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	applicant, err := c.experienceService.GetApplicantByID(userID)
	if err != nil {
		messError := fmt.Sprintf("failed to access create job experience, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	experience, err := c.experienceService.CreateExperience(experienceInput, int(applicant.ID))
	if err != nil {
		messError := fmt.Sprintf("failed to access create job experience, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "success to create job experience", experience)
	ctx.JSON(http.StatusOK, response)
}

func (c *experienceController) UpdateExperience(ctx *gin.Context) {
	var inputID dto.GetExperienceDetailDTO

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helpers.BuildErrorResponse("failed update experience", "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	var inputData dto.CreateExperienceDTO
	err = ctx.ShouldBind(&inputData)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helpers.BuildErrorResponse("failed update experience", "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		messError := fmt.Sprintf("failed to access update job experience, token user applicant wrong or empty")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		messError := fmt.Sprintf("failed to access update job experience, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	user, err := c.experienceService.GetUserByID(userID)
	if user.Role != "user" {
		messError := fmt.Sprintf("failed to update job experience, role is unknow")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	applicant, err := c.experienceService.GetApplicantByID(userID)
	if err != nil {
		messError := fmt.Sprintf("failed to update job experience, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	idApplicant := applicant.ID
	
	experience, err := c.experienceService.GetExperienceByID(int(idApplicant))
	if err != nil {
		messError := fmt.Sprintf("failed to update job experience, user %v %v have no job experience", applicant.FirstName, applicant.LastName)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	idApplicantExperience := experience.ApplicantID
	
	if idApplicant != idApplicantExperience {
		messError := fmt.Sprintf("failed to update job experience, not a user applicant owner experience")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	updateExperience, err := c.experienceService.UpdateExperience(inputID.ID, inputData, idApplicant)
	if err != nil {
		messError := fmt.Sprintf("failed to update job experience")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	message := fmt.Sprintf("success to update job experience")
	response := helpers.BuildResponse(true, message, updateExperience)
	ctx.JSON(http.StatusBadRequest, response)
	return
}

