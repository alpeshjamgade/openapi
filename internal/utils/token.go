package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"open-api/config"
	"open-api/internal/client"
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

	token.AccessUuid = GetUUID()
	token.RefreshUuid = GetUUID()

	var err error
	atClaims := jwt.MapClaims{
		"authorized":  true,
		"access_uuid": token.AccessUuid,
		"email":       email,
		"exp":         token.AtExpires,
		"user_type":   userType,
		"user_role":   userRole,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = accessToken.SignedString([]byte(jwtKey))
	if err != nil {
		Logger.Error(err)
		return "", err
	}

	cacheClient, err := client.GetCacheClient(ctx)
	if err != nil {
		Logger.Error(err)
		return "", err
	}

	expiry := time.Unix(token.AtExpires, 0)
	err = cacheClient.SaveAccessToken(ctx, email, token.AccessToken, expiry.Sub(currTime))
	return token.AccessToken, nil
}

func ValidateAuthToken(ctx context.Context, accessToken string) (*jwt.Token, error) {
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	jwtKey := config.JwtSecretKey
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		Logger.Error(err)
		return nil, errors.New("invalid token")
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, errors.New("invalid token")
	}

	fmt.Println(claims)
	email, ok := claims["email"].(string)
	if !ok {
		Logger.Error(err)
		return nil, errors.New("invalid token")
	}

	cacheClient, err := client.GetCacheClient(ctx)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}

	key, _ := cacheClient.GetAccessToken(ctx, email)
	if key != fmt.Sprintf("%s", accessToken) {
		Logger.Error("redis key not match")

		return nil, errors.New("invalid token")
	}

	return token, nil
}
