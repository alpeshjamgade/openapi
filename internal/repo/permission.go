package repo

import (
	"context"
	"openapi-client/internal/models"
	"strconv"
	"strings"
)

func (repo *Repo) CreatePermission(ctx context.Context, permission *models.Permission) error {
	_, err := repo.DB.DB().Exec(
		`INSERT INTO permissions(name) VALUES ($1)`,
		permission.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repo) GetPermissionByName(ctx context.Context, name string) (models.Permission, error) {

	var permission models.Permission
	sqlRow := repo.DB.DB().QueryRow(`SELECT id, name, created_at, updated_at FROM permissions WHERE name = $1`, name)
	err := sqlRow.Scan(&permission.ID, &permission.Name, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		return permission, err
	}
	return permission, nil
}
func (repo *Repo) UpdatePermission(ctx context.Context, permission *models.UpdatePermissionRequest) error {
	setClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if permission.NewName != "" {
		setClauses = append(setClauses, "name=$"+strconv.Itoa(argPos))
		args = append(args, permission.NewName)
		argPos++
	}

	if len(setClauses) == 0 {
		return nil
	}

	query := "UPDATE permissions SET " + strings.Join(setClauses, ", ") + " WHERE name=$" + strconv.Itoa(argPos)
	args = append(args, permission.Name)

	_, err := repo.DB.DB().Exec(query, args...)
	return err
}

func (repo *Repo) DeletePermissionByName(ctx context.Context, name string) error {

	_, err := repo.DB.DB().Exec(`DELETE FROM permissions WHERE name = $1`, name)

	return err
}
