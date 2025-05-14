package service

import (
	"context"
	"openapi-client/internal/models"
)

func (svc *Service) CreatePermission(ctx context.Context, permission *models.Permission) error {
	err := svc.repo.CreatePermission(ctx, permission)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) GetPermissionByName(ctx context.Context, email string) (models.Permission, error) {
	var permission models.Permission
	permission, err := svc.repo.GetPermissionByName(ctx, email)
	if err != nil {
		return permission, err
	}

	return permission, nil
}

func (svc *Service) UpdatePermission(ctx context.Context, permission *models.UpdatePermissionRequest) error {
	err := svc.repo.UpdatePermission(ctx, permission)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeletePermissionByName(ctx context.Context, email string) error {
	err := svc.repo.DeletePermissionByName(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
