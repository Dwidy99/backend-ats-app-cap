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

type EmployeeController interface {
	EditEmployee(ctx *gin.Context)
	FetchUserEmployee(ctx *gin.Context)
}

type employeeController struct {
	employeeService service.EmployeeService
	jwtService      service.JWTService
}

func NewEmployeeController(empService service.EmployeeService, jwtService service.JWTService) EmployeeController {
	return &employeeController{
		employeeService: empService,
		jwtService:      jwtService,
	}
}

func (c *employeeController) EditEmployee(ctx *gin.Context) {
	var employeeInput dto.EmployeeUpdateDTO

	err := ctx.ShouldBindJSON(&employeeInput)
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

	isAllowed, err := c.employeeService.IsAllowedToEdit(userID, uint64(id))
	if err != nil {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}

	if isAllowed {
		result, err := c.employeeService.UpdateEmployee(employeeInput, id)
		if err != nil {
			response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
			ctx.JSON(http.StatusForbidden, response)
		}
		response := helpers.BuildResponse(true, "ok", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}
}

func (c *employeeController) FetchUserEmployee(ctx *gin.Context) {
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

	result, err := c.employeeService.GetEmployeeById(userID)
	if err != nil {
		response := helpers.BuildErrorResponse("request failed", err.Error(), helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helpers.BuildResponse(true, "ok", result)
	ctx.JSON(http.StatusOK, response)
}
