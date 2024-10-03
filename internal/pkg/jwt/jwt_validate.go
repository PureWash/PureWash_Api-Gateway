package token

import (
	"api_gateway/internal/configs"
	"api_gateway/internal/pkg/logger"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(tokenStr string) (bool, error) {
	_, err := ExtractClaims(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}


func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	println(tokenStr)
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(configs.SignKey), nil
	})

	if err != nil {

		return nil, fmt.Errorf("faild to parse token: %w", err)
	}

	if !token.Valid {

		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {

		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
func GetUserIdByClaims(c *gin.Context, log logger.ILogger) (interface{}, error) {
	claimsRaw, exists := c.Get("claims")
	if !exists {
		log.Error("user not found")
		return nil, fmt.Errorf("user not found")
	}

	claims, ok := claimsRaw.(jwt.MapClaims)
	if !ok {
		log.Error("claims is not a valid map")
		return nil, fmt.Errorf("claims is not a valid map")
	}

	userId, ok := claims["id"].(string)
	if !ok {
		log.Error("user_id not found or is not a string")
		return nil, fmt.Errorf("user_id not found or is not a string")
	}
	return userId, nil
}
