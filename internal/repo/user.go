package repo

import (
	"context"
	"database/sql"
	"errors"
	"openapi-client/internal/models"
	"strconv"
	"strings"
)

func (repo *Repo) CreateUser(ctx context.Context, user *models.User) error {

	_, err := repo.DB.DB().Exec(
		`INSERT INTO users(name, email, phone, website, about, state, partner_id, password, role_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		user.Name,
		user.Email,
		user.Phone,
		user.Website,
		user.About,
		user.State,
		user.PartnerID,
		user.Password,
		user.RoleID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repo) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	sqlRow := repo.DB.DB().QueryRow(`SELECT id, name, email, phone, website, state, about, partner_id, password, role_id, created_at, updated_at FROM users WHERE email = $1`, email)

	err := sqlRow.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Website,
		&user.State,
		&user.About,
		&user.PartnerID,
		&user.Password,
		&user.RoleID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *Repo) UpdateUser(ctx context.Context, user *models.User) error {
	setClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if user.Name != "" {
		setClauses = append(setClauses, "name=$"+strconv.Itoa(argPos))
		args = append(args, user.Name)
		argPos++
	}
	if user.Email != "" {
		setClauses = append(setClauses, "email=$"+strconv.Itoa(argPos))
		args = append(args, user.Email)
		argPos++
	}
	if user.Phone != "" {
		setClauses = append(setClauses, "phone=$"+strconv.Itoa(argPos))
		args = append(args, user.Phone)
		argPos++
	}
	if user.Website != "" {
		setClauses = append(setClauses, "website=$"+strconv.Itoa(argPos))
		args = append(args, user.Website)
		argPos++
	}
	if user.About != "" {
		setClauses = append(setClauses, "about=$"+strconv.Itoa(argPos))
		args = append(args, user.About)
		argPos++
	}
	if user.State != "" {
		setClauses = append(setClauses, "state=$"+strconv.Itoa(argPos))
		args = append(args, user.State)
		argPos++
	}
	if user.PartnerID != "" {
		setClauses = append(setClauses, "partner_id=$"+strconv.Itoa(argPos))
		args = append(args, user.PartnerID)
		argPos++
	}
	if user.Password != "" {
		setClauses = append(setClauses, "password=$"+strconv.Itoa(argPos))
		args = append(args, user.Password)
		argPos++
	}

	if len(setClauses) == 0 {
		return nil
	}

	query := "UPDATE users SET " + strings.Join(setClauses, ", ") + " WHERE id=$" + strconv.Itoa(argPos)
	args = append(args, user.ID)

	_, err := repo.DB.DB().Exec(query, args...)
	return err
}

func (repo *Repo) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := repo.DB.DB().Exec(`DELETE FROM users WHERE email = $1`, email)
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repo) GetUserByEmailAndPassword(ctx context.Context, email, hashedPassword string) (models.User, error) {
	var user models.User
	sqlRow := repo.DB.DB().QueryRow(`SELECT 
    	id,
    	name,
    	email,
    	phone,
    	website,
    	about,
    	state,
    	partner_id,
    	password,
    	created_at,
    	updated_at
    FROM users WHERE email = $1 AND password =$2`, email, hashedPassword)

	err := sqlRow.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Phone,
		&user.Website,
		&user.About,
		&user.State,
		&user.PartnerID,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return user, errors.New("email or password is not correct")
	}
	if err != nil {
		return user, err
	}

	return user, nil
}
