package service

import (
	"context"
	"open-api-client/internal/models"
)

func (svc *Service) CreateUserApp(ctx context.Context, createUserAppRequest *models.CreateUserAppRequest) error {

	userApp := &models.UserApp{
		Name:          createUserAppRequest.Name,
		Type:          createUserAppRequest.Type,
		TradingID:     createUserAppRequest.TradingID,
		RedirectURL:   createUserAppRequest.RedirectURL,
		PostbackURL:   createUserAppRequest.PostbackURL,
		Description:   createUserAppRequest.Description,
		AppIconS3Path: createUserAppRequest.AppIconS3Path,
		UserID:        createUserAppRequest.UserID,
		PlanID:        createUserAppRequest.PlanID,
	}
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

func (svc *Service) UpdateUserApp(ctx context.Context, updateUserAppRequest *models.UpdateUserAppRequest) error {
	userApp := &models.UserApp{
		Name:          updateUserAppRequest.Name,
		Type:          updateUserAppRequest.Type,
		TradingID:     updateUserAppRequest.TradingID,
		RedirectURL:   updateUserAppRequest.RedirectURL,
		PostbackURL:   updateUserAppRequest.PostbackURL,
		Description:   updateUserAppRequest.Description,
		AppIconS3Path: updateUserAppRequest.AppIconS3Path,
		UserID:        updateUserAppRequest.UserID,
		PlanID:        updateUserAppRequest.PlanID,
	}

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
