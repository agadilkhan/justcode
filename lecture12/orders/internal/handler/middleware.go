package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"lecture12/orders/internal/config"
	"net/http"
	"strings"
)

func JWTVerify(cfg *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenString string
		tokenHeader := ctx.Request.Header.Get("authorization")
		tokenFields := strings.Fields(tokenHeader)
		if len(tokenFields) == 2 && tokenFields[0] == "Bearer" {
			tokenString = tokenFields[1]
		} else {
			ctx.AbortWithStatus(http.StatusForbidden)

			return
		}

		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("SgningMethodHMAC err: %v", token.Header["alg"])
			}

			return []byte(cfg.SecretKey), nil
		})
		if err != nil {
			ctx.AbortWithStatus(http.StatusForbidden)

			return
		}

		if !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		ctx.Next()
	}
}
