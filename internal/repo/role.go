package repo

import (
	"context"
	"openapi-client/internal/models"
	"strconv"
	"strings"
)

func (repo *Repo) CreateRole(ctx context.Context, role *models.Role) error {
	_, err := repo.DB.DB().Exec(
		`INSERT INTO roles(name) VALUES ($1)`,
		role.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repo) GetRoleByName(ctx context.Context, name string) (models.Role, error) {

	var role models.Role
	sqlRow := repo.DB.DB().QueryRow(`SELECT id, name, created_at, updated_at FROM roles WHERE name = $1`, name)
	err := sqlRow.Scan(&role.ID, &role.Name, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return role, err
	}
	return role, nil
}
func (repo *Repo) UpdateRole(ctx context.Context, role *models.UpdateRoleRequest) error {
	setClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if role.NewName != "" {
		setClauses = append(setClauses, "name=$"+strconv.Itoa(argPos))
		args = append(args, role.NewName)
		argPos++
	}

	if len(setClauses) == 0 {
		return nil
	}

	query := "UPDATE roles SET " + strings.Join(setClauses, ", ") + " WHERE name=$" + strconv.Itoa(argPos)
	args = append(args, role.Name)

	_, err := repo.DB.DB().Exec(query, args...)
	return err
}

func (repo *Repo) DeleteRoleByName(ctx context.Context, name string) error {

	_, err := repo.DB.DB().Exec(`DELETE FROM roles WHERE name = $1`, name)

	return err
}
