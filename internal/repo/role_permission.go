package repo

import (
	"context"
	"openapi-client/internal/models"
)

func (repo *Repo) CreateRolePermission(ctx context.Context, rolePermission *models.RolePermission) error {
	_, err := repo.DB.DB().Exec(
		`INSERT INTO role_permissions(role_id, permission_id) VALUES ($1, $2)`,
		rolePermission.RoleID, rolePermission.PermissionID,
	)

	if err != nil {
		return err
	}

	return nil
}
