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

type SkillController interface {
	CreateSkill(ctx *gin.Context)
	UpdateSkill(ctx *gin.Context)
	GetSkillByID(ctx *gin.Context)
	DeleteSkill(ctx *gin.Context)
	GetSkills(ctx *gin.Context)
}

type skillController struct {
	serviceSkill service.SkillService
	jwtService service.JWTService
}

func NewSkillController(skillService service.SkillService, jwtService service.JWTService) SkillController {
	return &skillController{
		serviceSkill: skillService,
		jwtService: jwtService,
	}
}

func (c *skillController) GetSkills(ctx *gin.Context) {
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
		messError := fmt.Sprintf("user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, err)
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

	applicant, err := c.serviceSkill.GetApplicantByID(userID)

	skills, err := c.serviceSkill.GetSkills(int(applicant.ID))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	response := helpers.BuildResponse(true, "success to get job skills", skills)
	ctx.JSON(http.StatusOK, response)
}

func (c *skillController) GetSkillByID(ctx *gin.Context) {
	var input dto.GetSkillDetailDTO

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

	if userID == 0 {
		errorMessage := gin.H{"error": errors.New("forbidden")}
		messError := fmt.Sprintf("user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, errorMessage)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	user, err := c.serviceSkill.GetUserByID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to update skill", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if user.Role != "user" {
		response := helpers.BuildErrorResponse("failed to process request", "role is not user", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	applicant, err := c.serviceSkill.GetApplicantByID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to get skill", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	// ambil data di tabel experience berdasarkan id url 
	skill, err := c.serviceSkill.GetSkillDetailByID(input.ID, int(applicant.ID))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", "failed to get job skill, id not found", helpers.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := helpers.BuildResponse(true, "job skill detail", skill)
	ctx.JSON(http.StatusOK, response)
}

func (c *skillController) CreateSkill(ctx *gin.Context) {
	var skillInput dto.Jobskill

	err := ctx.ShouldBind(&skillInput)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		messError := fmt.Sprintf("failed to create skill, token user applicant wrong or empty")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		messError := fmt.Sprintf("failed to access create job skill, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	user, err := c.serviceSkill.GetUserByID(userID)
	if user.Role != "user" {
		response := helpers.BuildErrorResponse("failed to process request", "role is not user", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	applicant, err := c.serviceSkill.GetApplicantByID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	experience, err := c.serviceSkill.CreateSkill(skillInput, int(applicant.ID))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "success to create job skill", experience)
	ctx.JSON(http.StatusOK, response)
}

func (c *skillController) UpdateSkill(ctx *gin.Context) {
	var inputID dto.GetSkillDetailDTO

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helpers.BuildErrorResponse("failed update skill", err.Error(), errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	var inputData dto.Jobskill
	err = ctx.ShouldBind(&inputData)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helpers.BuildErrorResponse("failed update skill", err.Error(), errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helpers.BuildErrorResponse("failed to access update job skill, token user applicant wrong or empty", errToken.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		response := helpers.BuildErrorResponse("failed to access request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	
	user, err := c.serviceSkill.GetUserByID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to access request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if user.Role != "user" {
		response := helpers.BuildErrorResponse("failed to process request", "role is not user", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	applicant, err := c.serviceSkill.GetApplicantByID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to update skill", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	skill, err := c.serviceSkill.UpdateSkill(inputID.ID, inputData, int(applicant.ID), userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to update skill", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "success to update job skill", skill)
	ctx.JSON(http.StatusOK, response)
}

func (c *skillController) DeleteSkill(ctx *gin.Context) {
	var inputID dto.GetSkillDetailDTO

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{"error": err}
		response := helpers.BuildErrorResponse("failed to delete skill", err.Error(), errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
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

	if userID == 0 {
		errorMessage := gin.H{"error": errors.New("forbidden")}
		messError := fmt.Sprintf("user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, errorMessage)
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	user, err := c.serviceSkill.GetUserByID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to update skill", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if user.Role != "user" {
		response := helpers.BuildErrorResponse("failed to process request", "role is not user", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	applicant, err := c.serviceSkill.GetApplicantByID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to update skill", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = c.serviceSkill.DeleteSkill(inputID.ID, int(applicant.ID))
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