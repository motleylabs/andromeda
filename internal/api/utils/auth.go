package utils

import (
	"andromeda/internal/api/models"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(data interface{}) (string, error) {
	token := jwt.New(jwt.GetSigningMethod(os.Getenv("JWT_SIGNING_METHOD")))

	claims := token.Claims.(jwt.MapClaims)
	if v, ok := data.(*models.User); ok {
		claims["ID"] = v.ID
	}

	tokenTimeout, err := strconv.ParseInt(os.Getenv("JWT_TOKEN_TIMEOUT"), 10, 64)
	if err != nil {
		return "", err
	}

	expire := time.Now().Add(time.Hour * time.Duration(tokenTimeout))

	claims["exp"] = expire.Unix()
	claims["orig_iat"] = time.Now().Unix()

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, err
}

func GenerateRefreshToken(data interface{}) (string, time.Time, error) {
	token := jwt.New(jwt.GetSigningMethod(os.Getenv("JWT_SIGNING_METHOD")))

	claims := token.Claims.(jwt.MapClaims)
	if v, ok := data.(*models.User); ok {
		claims["ID"] = v.ID
	}

	refreshTokenTimeout, err := strconv.ParseInt(os.Getenv("JWT_REFRESH_TOKEN_TIMEOUT"), 10, 64)
	if err != nil {
		return "", time.Time{}, err
	}

	expire := time.Now().Add(time.Hour * time.Duration(refreshTokenTimeout))

	claims["exp"] = expire.Unix()
	claims["orig_iat"] = time.Now().Unix()

	signedRefreshToken, err := token.SignedString([]byte(os.Getenv("JWT_REFRESH_TOKEN_SECRET_KEY")))
	if err != nil {
		return "", expire, err
	}

	return signedRefreshToken, expire, err
}

func GetClaimsFromJWT(c *gin.Context) (jwt.MapClaims, error) {
	token, err := ParseToken(c)
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{}
	for key, value := range token.Claims.(jwt.MapClaims) {
		claims[key] = value
	}

	return claims, nil
}

func ParseToken(c *gin.Context) (*jwt.Token, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		authHeader = c.Query("token")
	}
	if authHeader == "" {
		return nil, errors.New("auth header is empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return nil, errors.New("auth header is invalid")
	}

	token := parts[1]

	return jwt.Parse(token, func(jt *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(os.Getenv("JWT_SIGNING_METHOD")) != jt.Method {
			return nil, errors.New("invalid signing algorithm")
		}

		return []byte(os.Getenv("JWT_TOKEN_SECRET_KEY")), nil
	})
}
