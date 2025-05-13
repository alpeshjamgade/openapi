package service

import (
	"context"
	"openapi-client/internal/models"
	"openapi-client/internal/utils"
)

func (svc *Service) CreateUser(ctx context.Context, user *models.User) error {
	hashedPassword := utils.HashSHA256(user.Password)

	user.Password = string(hashedPassword)

	err := svc.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	user, err := svc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (svc *Service) UpdateUser(ctx context.Context, user *models.User) error {
	err := svc.repo.UpdateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteUserByEmail(ctx context.Context, email string) error {
	err := svc.repo.DeleteUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
