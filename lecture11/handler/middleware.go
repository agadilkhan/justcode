package handler

import (
	"github.com/gin-gonic/gin"
	"lecture8/service/jwttoken"
	"log"
	"net/http"
	"strings"
)

func JWTVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader("authorization")
		if authorizationHeader == "" {
			ctx.AbortWithStatus(http.StatusForbidden)
			log.Println("authorization header is not set")
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatus(http.StatusForbidden)
			log.Println("authorization header incorrect format")
			return
		}

		token := jwttoken.JWTToken{SecretKey: "lecture_8"}

		payload, err := token.ValidateToken(fields[1])
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			log.Println("invalid token")
			return
		}

		ctx.Set(authUserID, payload.UserID)
		ctx.Next()
	}
}
