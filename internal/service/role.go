package service

import (
	"context"
	"openapi-client/internal/models"
)

func (svc *Service) CreateRole(ctx context.Context, role *models.Role) error {
	err := svc.repo.CreateRole(ctx, role)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) GetRoleByName(ctx context.Context, email string) (models.Role, error) {
	var role models.Role
	role, err := svc.repo.GetRoleByName(ctx, email)
	if err != nil {
		return role, err
	}

	return role, nil
}

func (svc *Service) UpdateRole(ctx context.Context, role *models.UpdateRoleRequest) error {
	err := svc.repo.UpdateRole(ctx, role)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteRoleByName(ctx context.Context, email string) error {
	err := svc.repo.DeleteRoleByName(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
