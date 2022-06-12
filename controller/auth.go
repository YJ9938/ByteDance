package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/YJ9938/DouYin/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	jwt.StandardClaims
}

var secretKey = []byte(config.C.JWT.SecretKey)

// Auth middleware used to handle authentication for every route that needs auth
func AuthMiddleware(c *gin.Context) {
	// Check the token parameter
	token := c.Query("token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "鉴权参数错误",
		})
		return
	}

	// Check the token is valid and store user ID to the context
	claims := parseToken(token)
	if claims == nil {
		c.AbortWithStatusJSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "用户鉴权错误错误",
		})
		return
	}
	log.Printf("Token user ID: %s\n", claims.Id)

	// If auth success, we pass an 'id' to gin's context
	c.Set("id", claims.Id)
}

// signJWT signs a JWT and returns it.
func signJWT(userID int64) (string, error) {
	expireTime := time.Now().Add(time.Duration(config.C.JWT.ExpireMinutes) * time.Minute)
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			Id:        fmt.Sprintf("%d", userID),
			ExpiresAt: expireTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	return token.SignedString(secretKey)
}

// parseToken verify a JWT string and returns its claims.
func parseToken(tokenString string) *Claims {
	var claims Claims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil
	}
	return token.Claims.(*Claims)
}
