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

type ExperienceController interface {
	CreateExperience(ctx *gin.Context)
	UpdateExperience(ctx *gin.Context)
	DeleteExperience(ctx *gin.Context)
	GetAllExperiences(ctx *gin.Context)
	GetExperienceByID(ctx *gin.Context)
}

type experienceController struct {
	experienceService service.ExperienceService
	jwtService        service.JWTService
}

func NewExperienceController(expService service.ExperienceService, jwtService service.JWTService) ExperienceController {
	return &experienceController{
		experienceService: expService,
		jwtService:        jwtService,
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
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
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
		errMessage := fmt.Sprintf("role is %v, should role user to access the data", role)
		response := helpers.BuildErrorResponse("failed to process request", errMessage, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	applicant, err := c.experienceService.GetApplicantByID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	experience, err := c.experienceService.CreateExperience(experienceInput, int(applicant.ID))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
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
		response := helpers.BuildErrorResponse("failed update experience", err.Error(), errorMessage)
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
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
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
		errMessage := fmt.Sprintf("role is %v, should role user to access the data", role)
		response := helpers.BuildErrorResponse("failed to process request", errMessage, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil data di tabel experience berdasarkan id url
	experienceId, err := c.experienceService.GetExperienceByID(inputID.ID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
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
}

func (c *experienceController) DeleteExperience(ctx *gin.Context) {
	var inputID dto.GetExperienceDetailDTO

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helpers.BuildErrorResponse("failed to delete experience", err.Error(), errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
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
		errMessage := fmt.Sprintf("role is %v, should role user to access the data", role)
		response := helpers.BuildErrorResponse("failed to process request", errMessage, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil data di tabel experience berdasarkan id url
	experienceId, err := c.experienceService.GetExperienceByID(inputID.ID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", "user not found", helpers.EmptyObj{})
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
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	response := helpers.BuildResponse(true, "success to delete job experience", nil)
	ctx.JSON(http.StatusOK, response)
}

func (c *experienceController) GetAllExperiences(ctx *gin.Context) {

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
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	role := fmt.Sprintf("%v", claims["role"])
	if role != "user" || role == "" {
		errMessage := fmt.Sprintf("role is %v, should role user to access the data", role)
		response := helpers.BuildErrorResponse("failed to process request", errMessage, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
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

func (c *experienceController) GetExperienceByID(ctx *gin.Context) {
	var input dto.GetExperienceDetailDTO

	err := ctx.ShouldBindUri(&input)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

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
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	role := fmt.Sprintf("%v", claims["role"])
	if role != "user" || role == "" {
		errMessage := fmt.Sprintf("role is %v, should role user to access the data", role)
		response := helpers.BuildErrorResponse("failed to process request", errMessage, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil data di tabel experience berdasarkan id url
	experienceId, err := c.experienceService.GetExperienceByID(input.ID)
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

	experience, err := c.experienceService.GetExperienceByID(input.ID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "expereince detail", experience)
	ctx.JSON(http.StatusOK, response)
}
