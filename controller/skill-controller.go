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

type SkillController interface {
	CreateSkill(ctx *gin.Context)
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
	
	user, err := c.serviceSkill.GetUserByID(userID)
	if user.Role != "user" {
		messError := fmt.Sprintf("failed to access create job experience, role is unknow")
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	applicant, err := c.serviceSkill.GetApplicantByID(userID)
	if err != nil {
		messError := fmt.Sprintf("failed to access create job experience, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	experience, err := c.serviceSkill.CreateSkill(skillInput, int(applicant.ID))
	if err != nil {
		messError := fmt.Sprintf("failed to access create job experience, user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "success to create job experience", experience)
	ctx.JSON(http.StatusOK, response)
}