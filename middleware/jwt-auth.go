package middleware

import (
	"log"
	"mini-project/helpers"
	"mini-project/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := helpers.BuildErrorResponse("Failed to process request", "No token found", nil)
			ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			Claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", Claims["user_id"])
			log.Println("Claim[user_id]: ", Claims["issuers"])
		} else {
			log.Println(err)
			response := helpers.BuildErrorResponse("Token is not invalid", err.Error(), nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}