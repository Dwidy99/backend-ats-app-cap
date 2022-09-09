package controller

import (
	"fmt"
	"mini-project/dto"
	"mini-project/entity"
	"mini-project/helpers"
	"mini-project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthController interface {
	RegisterApplicants(ctx *gin.Context)
	RegisterEmployees(ctx *gin.Context)
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService service.JWTService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) RegisterApplicants(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO

	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := helpers.BuildErrorResponse("Failed to process request", "Duplicate email", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.ID, 10))
		createdUser.Token = token
		response := helpers.BuildResponse(true, "ok", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

func (c *authController) RegisterEmployees(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	role := fmt.Sprintf("%v", claims["role"])

	var registerDTO dto.RegisterEmployeeDTO

	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if role == "superadmin" {
		if !c.authService.IsDuplicateEmail(registerDTO.Email) {
			response := helpers.BuildErrorResponse("Email is registered", "Failed to process request", helpers.EmptyObj{})
			ctx.JSON(http.StatusConflict, response)
		} else {
			createdUser := c.authService.CreateEmployee(registerDTO)
			response := helpers.BuildResponse(true, "ok", createdUser)
			ctx.JSON(http.StatusCreated, response)
		}
	}

}

func (c *authController) Login(ctx *gin.Context) {
	var login dto.LoginDTO

	errDTO := ctx.ShouldBind(&login)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authResult := c.authService.VerifyCredential(login.Email, login.Password)
	if v, ok := authResult.(entity.User); ok {
		GeneratedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = GeneratedToken
		response := helpers.BuildResponse(true, "ok", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := helpers.BuildErrorResponse("Please check again your credential", "Invalid Credential", helpers.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}