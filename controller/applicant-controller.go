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

type ApplicantController interface {
	EditApplicant(ctx *gin.Context)
}

type applicantController struct {
	applicantService service.ApplicantService
	jwtService  service.JWTService
}

// NewAuthController creates a new instance of AuthController
func NewApplicantController(applicantService service.ApplicantService, jwtService service.JWTService) ApplicantController {
	return &applicantController{
		applicantService: applicantService,
		jwtService:  jwtService,
	}
}

func (c *applicantController) EditApplicant(ctx *gin.Context) {
	var inputID dto.ApplicantDTO
	
	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	
	var inputData dto.ApplicantUpdateDTO
	
	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(errToken.Error())
	}

	if userID != 0 {
		result := c.applicantService.UpdateApplicant(inputID, inputData)
		response := helpers.BuildResponse(true, "ok", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		res := helpers.BuildErrorResponse("Failed to process request", "error", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}
}