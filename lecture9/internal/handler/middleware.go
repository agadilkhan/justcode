package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"lecture9/internal/service/jwttoken"
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

const authUserID = "userID"

func getAuthUserID(ctx *gin.Context) (int, error) {
	userID, ok := ctx.MustGet(authUserID).(int)
	if !ok {
		ctx.JSON(http.StatusBadRequest, Error{
			Message: "can't get user id from auth",
		})
		return 0, errors.New("can't get user id from auth")
	}

	return userID, nil
}
