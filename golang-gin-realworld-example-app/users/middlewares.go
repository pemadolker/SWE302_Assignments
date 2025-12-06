package users

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"realworld-backend/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"errors"
)

// Strips 'TOKEN ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

// Extract token from Authorization header
func extractTokenFromRequest(c *gin.Context) (string, error) {
	// Try Authorization header first
	authHeader := c.GetHeader("Authorization")
	if authHeader != "" {
		return stripBearerPrefixFromTokenString(authHeader)
	}
	
	// Try query parameter
	tokenParam := c.Query("access_token")
	if tokenParam != "" {
		return tokenParam, nil
	}
	
	return "", errors.New("no token found")
}

// A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, my_user_id uint) {
	var myUserModel UserModel
	if my_user_id != 0 {
		db := common.GetDB()
		db.First(&myUserModel, my_user_id)
	}
	c.Set("my_user_id", my_user_id)
	c.Set("my_user_model", myUserModel)
}

// You can custom middlewares yourself as the doc: https://github.com/gin-gonic/gin#custom-middleware
//  r.Use(AuthMiddleware(true))
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUserModel(c, 0)
		
		// Extract token from request
		tokenString, err := extractTokenFromRequest(c)
		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		
		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(common.NBSecretPassword), nil
		})
		
		if err != nil {
			if auto401 {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		
		// Extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Validate required claims
			if id, exists := claims["id"]; exists {
				my_user_id := uint(id.(float64))
				UpdateContextUserModel(c, my_user_id)
			} else if auto401 {
				c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token claims"))
			}
		} else if auto401 {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		}
	}
}