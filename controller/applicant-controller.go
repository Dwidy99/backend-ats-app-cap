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
	var applicantUpdateDTO dto.ApplicantUpdateDTO
	var applicant dto.ApplicantDTO

	errDTO := ctx.ShouldBind(&applicantUpdateDTO)
	if errDTO != nil {
		res := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
	}

	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	fmt.Println("USERID: ", userID)
	fmt.Println("APPLICANT UserID: ", applicant.UserID)
	if c.applicantService.IsAllowedToEdit(userID, applicant.UserID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			applicantUpdateDTO.UserID = id
		}
		result := c.applicantService.UpdateApplicant(applicantUpdateDTO)
		response := helpers.BuildResponse(true, "ok", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helpers.BuildErrorResponse("update user failed", "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
	}
}