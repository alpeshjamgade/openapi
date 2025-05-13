package service

import (
	"context"
	"errors"
	"open-api-client/internal/logger"
	"open-api-client/internal/models"
	"open-api-client/internal/utils"
)

func (svc *Service) Login(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error) {

	hashedPassword := utils.HashSHA256(loginRequest.Password)

	user, err := svc.repo.GetUserByEmailAndPassword(ctx, loginRequest.Email, string(hashedPassword))
	if err != nil {
		return nil, err
	} else if user.Email == "" {
		return nil, errors.New("user not found")
	}

	authToken, err := utils.GenerateAuthToken(ctx, loginRequest.Email, "user", "user")
	if err != nil {
		return nil, err
	}

	loginResponse := &models.LoginResponse{Email: loginRequest.Email, AuthToken: authToken}

	return loginResponse, nil
}

func (svc *Service) AdminLogin(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error) {

	Logger := logger.CreateFileLoggerWithCtx(ctx)

	hashedPassword := utils.HashSHA256(loginRequest.Password)

	user, err := svc.repo.GetAdminByEmailAndPassword(ctx, loginRequest.Email, string(hashedPassword))
	if err != nil {
		Logger.Errorw("error while fetching admin", "error", err)

		return nil, err
	} else if user.Email == "" {
		return nil, errors.New("user not found")
	}

	authToken, err := utils.GenerateAuthToken(ctx, loginRequest.Email, "admin", "admin")
	if err != nil {
		Logger.Errorw("generate auth token error", "error", err)
		return nil, err
	}

	loginResponse := &models.LoginResponse{Email: loginRequest.Email, AuthToken: authToken}

	return loginResponse, nil
}
