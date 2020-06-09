package auth

import (
	"github.com/dgrijalva/jwt-go"
	"nowim.user/internal/config"
	"strconv"
	"time"
)

func GenerateToken(userID int64, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   strconv.FormatInt(userID, 10),
		"username": username,
		"time":     time.Now().String(),
	})
	return token.SignedString([]byte(config.Config().Secret))
}
