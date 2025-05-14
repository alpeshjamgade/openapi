package utils

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"openapi-client/config"
	"openapi-client/internal/client"
	"openapi-client/internal/logger"
	"openapi-client/internal/models"
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

func GenerateAuthToken(ctx context.Context, id int, name string, email string, userType string) (string, error) {

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
		"id":          id,
		"email":       email,
		"name":        name,
		"exp":         token.AtExpires,
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
	cacheKey := AccessTokenKey(email, userType)
	err = cacheClient.SaveAccessToken(ctx, cacheKey, token.AccessToken, expiry.Sub(currTime))
	return token.AccessToken, nil
}

func ValidateAuthToken(ctx context.Context, accessToken string, userType string) (*models.Session, error) {
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

	idFloat, ok := claims["id"].(float64)
	if !ok {
		return nil, errors.New("invalid token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	name, ok := claims["name"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	session := &models.Session{
		UserID: int(idFloat),
		Email:  email,
		Name:   name,
	}

	cacheClient, err := client.GetCacheClient(ctx)
	if err != nil {
		Logger.Error(err)
		return nil, err
	}

	cacheKey := AccessTokenKey(email, userType)
	cachedToken, _ := cacheClient.GetAccessToken(ctx, cacheKey)
	if cachedToken != fmt.Sprintf("%s", accessToken) {
		Logger.Error("redis key not match")

		return nil, errors.New("invalid token")
	}

	return session, nil
}

func AccessTokenKey(email string, userType string) string {
	return fmt.Sprintf("token:%s:%s", userType, email)
}
