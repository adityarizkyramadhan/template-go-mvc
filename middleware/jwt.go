package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/adityarizkyramadhan/template-go-mvc/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(roles []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the Authorization header
		token := ctx.GetHeader("Authorization")
		if token == "" {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, errors.New("token not found").Error())
			ctx.Abort()
			return
		}
		token = strings.Replace(token, "Bearer ", "", 1)
		claims, err := verifyJWT(token)
		if err != nil {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			ctx.Abort()
			return
		}
		// Check if the role in the token matches one of the desired roles
		mapClaims, ok := claims.(jwt.MapClaims)
		if !ok {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "invalid token claims")
			ctx.Abort()
			return
		}
		role, ok := mapClaims["role"].(string)
		if !ok {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "role not found in token")
			ctx.Abort()
			return
		}
		roleAllowed := false
		for _, r := range roles {
			if r == role {
				roleAllowed = true
				break
			}
		}
		if !roleAllowed {
			utils.ErrorResponse(ctx, http.StatusUnauthorized, "role not allowed")
			ctx.Abort()
			return
		}
		ctx.Set("id", mapClaims["id"])
		ctx.Set("role", role)
		ctx.Set("email", mapClaims["username"])
		ctx.Next()
	}
}

func verifyJWT(tokenString string) (interface{}, error) {
	var secretKey = os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("SECRET_KEY environment variable not set")
	}
	var jwtSecret = []byte(secretKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Invalid token")
}
