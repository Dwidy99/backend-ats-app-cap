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
	FetchUserApplicant(ctx *gin.Context)
	UploadAvatar(ctx *gin.Context)
	DetailApplicant(ctx *gin.Context)
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

func (c *applicantController) DetailApplicant(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	applicant, err := c.applicantService.GetDetailApplicant(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	res := helpers.BuildResponse(true, "successfuly get data user applicant", applicant)
	ctx.JSON(http.StatusOK, res)
}

func (c *applicantController) EditApplicant(ctx *gin.Context) {
	var applicantInput dto.ApplicantUpdateDTO

	err := ctx.ShouldBindJSON(&applicantInput)
	if err != nil {
		res := helpers.BuildErrorResponse("Failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	id, _ := strconv.Atoi(ctx.Param("id"))

	isAllowed, err := c.applicantService.IsAllowedToEdit(userID, uint64(id))
	if err != nil {
		response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	if isAllowed {
		result, err := c.applicantService.UpdateApplicant(applicantInput, id)
		if err != nil {
			response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		response := helpers.BuildResponse(true, "ok", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}
}

func (c *applicantController) FetchUserApplicant(ctx *gin.Context) {

	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	applicant, err := c.applicantService.GetApplicantByUserID(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	res := helpers.BuildResponse(true, "successfuly get data user applicant", applicant)
	ctx.JSON(http.StatusOK, res)
}

func (c *applicantController) UploadAvatar(ctx *gin.Context) {
	file, err := ctx.FormFile("avatar")
	if err != nil {
		response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}
	claims := token.Claims.(jwt.MapClaims)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", claims["user_id"]))
	if err != nil {
		messError := fmt.Sprintf("user applicant with user id %v is empty", userID)
		response := helpers.BuildErrorResponse("failed to process request", messError, helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
		return
	}

	path := fmt.Sprintf("images/applicants/%d-%s", userID, file.Filename)
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = c.applicantService.UploadAvatar(userID, path)
	if err != nil {
		response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	res := helpers.BuildResponse(true, "successfuly uploaded avatar", data)
	ctx.JSON(http.StatusOK, res)
}