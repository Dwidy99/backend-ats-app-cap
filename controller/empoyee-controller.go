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

type EmployeeController interface {
	EditEmployee(ctx *gin.Context)
	FetchUserEmployee(ctx *gin.Context)
}

type employeeController struct {
	employeeService service.EmployeeService
	jwtService service.JWTService
}

func NewEmployeeController(employeeService service.EmployeeService, jwtService service.JWTService) EmployeeController {
	return &employeeController{
		employeeService: employeeService,
		jwtService: jwtService,
	}
}

func (c *employeeController) EditEmployee(ctx *gin.Context)  {
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

	if c.employeeService.IsAllowedToEdit(userID, uint64(id)) {
		
		result := c.employeeService.UpdateEmployee(employeeInput, id)
		response := helpers.BuildResponse(true, "ok", result)
		ctx.JSON(http.StatusOK, response)
	} else {
		response := helpers.BuildErrorResponse("You dont have permission", "You are not the owner", helpers.EmptyObj{})
		ctx.JSON(http.StatusForbidden, response)
	}
}

func (c *employeeController) FetchUserEmployee(ctx *gin.Context) {
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
	
	result := c.employeeService.GetEmployeeById(userID)
	response := helpers.BuildResponse(true, "ok", result)
	ctx.JSON(http.StatusOK, response)
}