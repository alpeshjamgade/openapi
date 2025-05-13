package repo

import (
	"context"
	"database/sql"
	"errors"
	"open-api-client/internal/models"
	"strconv"
	"strings"
)

func (repo *Repo) CreateAdmin(ctx context.Context, admin *models.Admin) error {

	_, err := repo.DB.DB().Exec(
		`INSERT INTO admins(name, email, phone, status, password) VALUES($1, $2, $3, $4, $5)`,
		admin.Name,
		admin.Email,
		admin.Phone,
		"active",
		admin.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repo) GetAdminByEmail(ctx context.Context, email string) (models.Admin, error) {
	var admin models.Admin
	sqlRow := repo.DB.DB().QueryRow(`SELECT 
    	id,
		name, 
		email, 
		phone, 
		status, 
		created_at, 
		updated_at, 
		last_login
    FROM admins WHERE email = $1`, email)

	err := sqlRow.Scan(
		&admin.ID,
		&admin.Name,
		&admin.Email,
		&admin.Phone,
		&admin.Status,
		&admin.CreatedAt,
		&admin.UpdatedAt,
		&admin.LastLogin,
	)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (repo *Repo) UpdateAdmin(ctx context.Context, admin *models.Admin) error {
	setClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if admin.Name != "" {
		setClauses = append(setClauses, "name=$"+strconv.Itoa(argPos))
		args = append(args, admin.Name)
		argPos++
	}
	if admin.Email != "" {
		setClauses = append(setClauses, "email=$"+strconv.Itoa(argPos))
		args = append(args, admin.Email)
		argPos++
	}
	if admin.Phone != "" {
		setClauses = append(setClauses, "phone=$"+strconv.Itoa(argPos))
		args = append(args, admin.Phone)
		argPos++
	}
	if admin.Status != "" {
		setClauses = append(setClauses, "partner_id=$"+strconv.Itoa(argPos))
		args = append(args, admin.Status)
		argPos++
	}
	if admin.Password != "" {
		setClauses = append(setClauses, "password=$"+strconv.Itoa(argPos))
		args = append(args, admin.Password)
		argPos++
	}

	if len(setClauses) == 0 {
		return nil
	}

	query := "UPDATE admins SET " + strings.Join(setClauses, ", ") + " WHERE id=$" + strconv.Itoa(argPos)
	args = append(args, admin.ID)

	_, err := repo.DB.DB().Exec(query, args...)
	return err
}

func (repo *Repo) DeleteAdminByEmail(ctx context.Context, email string) error {
	_, err := repo.DB.DB().Exec(`DELETE FROM admins WHERE email = $1`, email)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repo) GetAdminByEmailAndPassword(ctx context.Context, email, hashedPassword string) (models.Admin, error) {
	var admin models.Admin
	sqlRow := repo.DB.DB().QueryRow(`SELECT 
    	id,
    	name,
    	email,
    	phone,
    	status,
    	created_at,
    	updated_at,
    	last_login
    FROM admins WHERE email = $1 AND password =$2`, email, hashedPassword)

	err := sqlRow.Scan(&admin.ID, &admin.Name, &admin.Email, &admin.Phone, &admin.Status, &admin.CreatedAt, &admin.UpdatedAt, &admin.LastLogin)
	if errors.Is(err, sql.ErrNoRows) {
		return admin, errors.New("email or password is not correct")
	}
	if err != nil {
		return admin, err
	}

	return admin, nil
}
