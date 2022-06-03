package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yamess/auth/configs"
	"net/http"
	"strings"
)

func Authenticate(g *gin.Context) {
	headers := g.GetHeader("Authorization")
	tokenString := strings.Split(headers, " ")[1]

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.SecretKey), nil
		},
	)

	if err != nil {
		g.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	g.Set("claims", claims)
	g.Next()
}
