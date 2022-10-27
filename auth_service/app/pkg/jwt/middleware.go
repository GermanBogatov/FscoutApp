package jwt

import (
	"errors"
	"fmt"
	"github.com/GermanBogatov/auth_service/internal/config"
	"github.com/GermanBogatov/auth_service/pkg/logging"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

const (
	sportsman = "sportsman"
	scout     = "scout"
	admin     = "admin"
)

func MiddlewareSportsman() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logging.GetLogger()
		authHeader := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			logger.Error("Malformed token")
			c.Writer.WriteHeader(http.StatusUnauthorized)
			c.Writer.Write([]byte("Malformed token"))
			return
		}

		accessToken := authHeader[1]
		key := []byte(config.GetConfig().JWT.Secret)

		logger.Debug("Parsing jwt-token")
		token, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return key, nil
		})
		if err != nil {
			unauthorized(c.Writer, err)
			return
		}
		fmt.Println("token: ", token.Claims.(*UserClaims).Audience)
		if !token.Valid {
			logger.Error("token has been inspired")
			unauthorized(c.Writer, err)
			return
		}
		claims, ok := token.Claims.(*UserClaims)
		if !ok {
			unauthorized(c.Writer, err)
			return
		}
		if claims.Audience != sportsman {
			logger.Error("role does not match")
			unauthorized(c.Writer, err)
		}
		c.Next()

	}
}

func unauthorized(w http.ResponseWriter, err error) {
	logging.GetLogger().Error(err)
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("unauthorized"))
}
