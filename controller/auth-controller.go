package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PutraFajarF/backend-ats-app-cap/dto"
	"github.com/PutraFajarF/backend-ats-app-cap/entity"
	"github.com/PutraFajarF/backend-ats-app-cap/helpers"
	"github.com/PutraFajarF/backend-ats-app-cap/service"

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
	jwtService  service.JWTService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) RegisterEmployees(ctx *gin.Context) {
	var RegisterEmployeeDTO dto.RegisterEmployeeDTO

	errEmployeeDTO := ctx.ShouldBind(&RegisterEmployeeDTO)
	if errEmployeeDTO != nil {
		resp := helpers.BuildErrorResponse("Failed to process request", errEmployeeDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	if !c.authService.IsDuplicateEmail(RegisterEmployeeDTO.Email) {
		resp := helpers.BuildErrorResponse("Failed to process request", "Duplicate email", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, resp)
		return
	} else {
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
			response := helpers.BuildErrorResponse("failed to process request", "user not login", helpers.EmptyObj{})
			ctx.JSON(http.StatusForbidden, response)
			return
		}

		user, err := c.authService.GetUserByID(userID)
		if err != nil {
			response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		if user.Role != "admin" {
			response := helpers.BuildErrorResponse("request blocked, role invalid", "role not superadmin", helpers.EmptyObj{})
			ctx.JSON(http.StatusForbidden, response)
			return
		} else {
			createdUser, err := c.authService.CreateEmployee(RegisterEmployeeDTO)
			if err != nil {
				response := helpers.BuildErrorResponse("failed to process request", err.Error(), helpers.EmptyObj{})
				ctx.JSON(http.StatusBadRequest, response)
				return
			}
			response := helpers.BuildResponse(true, "ok", createdUser)
			ctx.JSON(http.StatusCreated, response)
		}
	}

}

func (c *authController) RegisterApplicants(ctx *gin.Context) {
	var registerApplicantDTO dto.RegisterApplicantDTO

	errDTO := ctx.ShouldBind(&registerApplicantDTO)
	if errDTO != nil {
		response := helpers.BuildErrorResponse("Failed to process request", errDTO.Error(), helpers.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerApplicantDTO.Email) {
		response := helpers.BuildErrorResponse("Failed to process request", "Duplicate email", helpers.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateApplicant(registerApplicantDTO)
		response := helpers.BuildResponse(true, "ok", createdUser)
		ctx.JSON(http.StatusCreated, response)
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
