package middleware

import (
	"App/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var Uid uint64

func AuthoriseJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BEARER = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		tokenStr := authHeader[len(BEARER):]

		token, err := service.NewJWTService().Validate(tokenStr)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("ID: ", claims["id"])
			log.Println("Expiration ", claims["exp"])
			Uid = uint64(claims["id"].(float64))
		} else {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
