package service

import (
	"context"
	"open-api/internal/models"
)

func (svc *Service) CreateUserApp(ctx context.Context, userApp *models.UserApp) error {

	err := svc.repo.CreateUserApp(ctx, userApp)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) GetUserAppByID(ctx context.Context, id string) (models.UserApp, error) {
	var userApp models.UserApp
	userApp, err := svc.repo.GetUserAppByID(ctx, id)
	if err != nil {
		return userApp, err
	}

	return userApp, nil
}

func (svc *Service) UpdateUserApp(ctx context.Context, userApp *models.UserApp) error {
	err := svc.repo.UpdateUserApp(ctx, userApp)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteUserAppByID(ctx context.Context, id string) error {
	err := svc.repo.DeleteUserAppByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
