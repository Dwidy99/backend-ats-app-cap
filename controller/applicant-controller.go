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
	jwtService       service.JWTService
}

// NewAuthController creates a new instance of AuthController
func NewApplicantController(applicantService service.ApplicantService, jwtService service.JWTService) ApplicantController {
	return &applicantController{
		applicantService: applicantService,
		jwtService:       jwtService,
	}
}

func (c *applicantController) EditApplicant(ctx *gin.Context) {
	var applicantUpdateDTO dto.ApplicantUpdateDTO

	err := ctx.ShouldBindJSON(&applicantUpdateDTO)
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
	userID := fmt.Sprintf("%v", claims["user_id"])
	id, _ := strconv.Atoi(ctx.Param("id"))

	if c.applicantService.IsAllowedToEdit(userID, uint64(id)) {
		
		result := c.applicantService.UpdateApplicant(applicantUpdateDTO)
		response := helpers.BuildResponse(true, "OK", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}
}
