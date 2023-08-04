package helper

import (
	"errors"
	"hotelPMS/auth"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			AbortMessage(c, errors.New("unauthorized"))
			return
		}

		tokenString := ""
		tokenList := strings.Split(authHeader, " ")
		if len(tokenList) == 2 {
			tokenString = tokenList[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			AbortMessage(c, errors.New("Unauthorized"))
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			AbortMessage(c, errors.New("Unauthorized"))
			return
		}

		userID := int(claim["user_id"].(float64))
		err = GetUser(c, userID)
		if err != nil {
			AbortMessage(c, errors.New("Unauthorized"))
			return
		}
	}
}
