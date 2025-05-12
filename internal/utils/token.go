package utils

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"open-api/config"
	"open-api/internal/logger"
	"time"
)

type JWTToken struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func GenerateAuthToken(ctx context.Context, email string, userType string, userRole string) (string, error) {

	Logger := logger.CreateFileLoggerWithCtx(ctx)
	jwtKey := config.JwtSecretKey

	token := &JWTToken{}
	currTime := time.Now()
	token.AtExpires = currTime.Add(time.Hour * 24).Unix()
	token.RtExpires = currTime.Add(time.Hour * 24 * 7).Unix()

	token.AccessUuid = GetUUID()
	token.RefreshUuid = GetUUID()

	var err error
	atClaims := jwt.MapClaims{
		"authorized":  true,
		"access_uuid": true,
		"user_name":   true,
		"exp":         true,
		"user_type":   true,
		"user_role":   true,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = accessToken.SignedString([]byte(jwtKey))
	if err != nil {
		Logger.Error(err)
		return "", err
	}

	return token.AccessToken, nil
}
