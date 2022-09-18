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

type ExperienceController interface {
	CreateExperience(ctx *gin.Context)
	UpdateExperience(ctx *gin.Context)
	DeleteExperience(ctx *gin.Context)
	GetAllExperiences(ctx *gin.Context)
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

	// ambil data di tabel experience berdasarkan id url 
	experienceId, err := c.experienceService.GetExperienceByID(inputID.ID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", "failed to update job experience, id not found", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	idApplicant := experienceId.ApplicantID

	// ambil data di tabel applicant berdasarkan id yang login
	applicant, _ := c.experienceService.GetApplicantByID(userID)
	// ambil data di tabel experience berdasarkan id applicant
	experienceApplicant, err := c.experienceService.GetExperienceByIdApplicant(int(applicant.ID))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// cek apakah applicant_id (yg ada di tabel jobexperience) tidak sama dengan applicant_id (yg ada di tabel applicant)
	if idApplicant != experienceApplicant.ApplicantID {
		response := helpers.BuildErrorResponse("failed to process request", "not an owner this job experience", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	updateExperience, err := c.experienceService.UpdateExperience(inputID.ID, inputData)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "success to update job experience", updateExperience)
	ctx.JSON(http.StatusOK, response)
	return
}

func (c *experienceController) DeleteExperience(ctx *gin.Context) {
	var inputID dto.GetExperienceDetailDTO

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helpers.BuildErrorResponse("failed to delete experience", "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		messError := fmt.Sprintf("failed to access delete job experience, token user applicant wrong or empty")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		messError := fmt.Sprintf("failed to delete job experience, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if userID == 0 {
		errorMessage := gin.H{"error": errors.New("forbidden")}
		messError := fmt.Sprintf("failed to delete job experience, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, errorMessage)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	// ambil data di tabel experience berdasarkan id url 
	experienceId, err := c.experienceService.GetExperienceByID(inputID.ID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", "failed to update job experience, id not found", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	idApplicant := experienceId.ApplicantID

	// ambil data di tabel applicant berdasarkan id yang login
	applicant, _ := c.experienceService.GetApplicantByID(userID)
	// ambil data di tabel experience berdasarkan id applicant
	experienceApplicant, err := c.experienceService.GetExperienceByIdApplicant(int(applicant.ID))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// cek apakah applicant_id (yg ada di tabel jobexperience) tidak sama dengan applicant_id (yg ada di tabel applicant)
	if idApplicant != experienceApplicant.ApplicantID {
		response := helpers.BuildErrorResponse("failed to process request", "not an owner this job experience", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = c.experienceService.DeleteExperience(inputID.ID)
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		messError := fmt.Sprintf("failed to delete job experience")
		response := helpers.BuildErrorResponse("failed to process request", messError, errorMessage)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	response := helpers.BuildResponse(true, "success to delete job experience", nil)
	ctx.JSON(http.StatusOK, response)
}

func (c *experienceController) GetAllExperiences(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		messError := fmt.Sprintf("failed to access delete job experience, token user applicant wrong or empty")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		messError := fmt.Sprintf("failed to delete job experience, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if userID == 0 {
		errorMessage := gin.H{"error": errors.New("forbidden")}
		messError := fmt.Sprintf("failed to delete job experience, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, errorMessage)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	experience, err := c.experienceService.GetAllExperiences(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	response := helpers.BuildResponse(true, "success to get job experiences", experience)
	ctx.JSON(http.StatusOK, response)
}