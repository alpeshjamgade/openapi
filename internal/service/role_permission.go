package service

import (
	"context"
	"openapi-client/internal/models"
)

func (svc *Service) CreateRolePermission(ctx context.Context, rolePermission *models.RolePermission) error {
	err := svc.repo.CreateRolePermission(ctx, rolePermission)
	if err != nil {
		return err
	}

	return nil
}
