package service

import (
	"context"
	"errors"
	"openapi-client/internal/constants"
	"openapi-client/internal/logger"
	"openapi-client/internal/models"
	"openapi-client/internal/utils"
)

func (svc *Service) Login(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error) {

	hashedPassword := utils.HashSHA256(loginRequest.Password)

	user, err := svc.repo.GetUserByEmailAndPassword(ctx, loginRequest.Email, string(hashedPassword))
	if err != nil {
		return nil, err
	} else if user.Email == "" {
		return nil, errors.New("user not found")
	}

	authToken, err := utils.GenerateAuthToken(ctx, user.ID, user.Name, user.Email, constants.Client)
	if err != nil {
		return nil, err
	}

	loginResponse := &models.LoginResponse{Email: loginRequest.Email, AuthToken: authToken, UserID: user.ID}

	return loginResponse, nil
}

func (svc *Service) AdminLogin(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error) {

	Logger := logger.CreateFileLoggerWithCtx(ctx)

	hashedPassword := utils.HashSHA256(loginRequest.Password)

	admin, err := svc.repo.GetAdminByEmailAndPassword(ctx, loginRequest.Email, string(hashedPassword))
	if err != nil {
		Logger.Errorw("error while fetching admin", "error", err)

		return nil, err
	} else if admin.Email == "" {
		return nil, errors.New("user not found")
	}

	authToken, err := utils.GenerateAuthToken(ctx, admin.ID, admin.Name, admin.Email, constants.Admin)
	if err != nil {
		Logger.Errorw("generate auth token error", "error", err)
		return nil, err
	}

	loginResponse := &models.LoginResponse{Email: loginRequest.Email, AuthToken: authToken, UserID: admin.ID}

	return loginResponse, nil
}
