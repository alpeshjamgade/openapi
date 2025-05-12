package repo

import (
	"context"
	"open-api/internal/models"
	"strconv"
	"strings"
)

func (repo *Repo) CreateUserApp(ctx context.Context, userApp *models.UserApp) error {

	_, err := repo.DB.DB().Exec(
		`INSERT INTO user_apps(name, trading_id, redirect_url, postback_url, description, app_icon_s3_path, user_id, plan_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		userApp.Name,
		userApp.TradingID,
		userApp.RedirectURL,
		userApp.PostbackURL,
		userApp.Description,
		userApp.AppIconS3Path,
		userApp.UserID,
		userApp.PlanID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repo) GetUserAppByID(ctx context.Context, id string) (models.UserApp, error) {
	var userApp models.UserApp
	sqlRow := repo.DB.DB().QueryRow(`SELECT *
    	id,
    	name,
    	type,
    	trading_id,
    	redirect_url,
    	postback_url,
    	description,
    	app_icon_s3_path,
    	user_id,
    	plan_id,
    	created_at,
    	updated_at,
    FROM user_apps WHERE id = $1`, id)

	err := sqlRow.Scan(
		&userApp.ID,
		&userApp.Name,
		&userApp.Type,
		&userApp.TradingID,
		&userApp.RedirectURL,
		&userApp.PostbackURL,
		&userApp.Description,
		&userApp.AppIconS3Path,
		&userApp.UserID,
		&userApp.PlanID,
		&userApp.CreatedAt,
		&userApp.UpdatedAt,
	)
	if err != nil {
		return userApp, err
	}

	return userApp, nil
}

func (repo *Repo) UpdateUserApp(ctx context.Context, userApp *models.UserApp) error {
	setClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if userApp.Name != "" {
		setClauses = append(setClauses, "name=$"+strconv.Itoa(argPos))
		args = append(args, userApp.Name)
		argPos++
	}
	if userApp.TradingID != "" {
		setClauses = append(setClauses, "id=$"+strconv.Itoa(argPos))
		args = append(args, userApp.ID)
		argPos++
	}
	if userApp.RedirectURL != "" {
		setClauses = append(setClauses, "phone=$"+strconv.Itoa(argPos))
		args = append(args, userApp.RedirectURL)
		argPos++
	}
	if userApp.PostbackURL != "" {
		setClauses = append(setClauses, "partner_id=$"+strconv.Itoa(argPos))
		args = append(args, userApp.PostbackURL)
		argPos++
	}
	if userApp.Description != "" {
		setClauses = append(setClauses, "password=$"+strconv.Itoa(argPos))
		args = append(args, userApp.Description)
		argPos++
	}
	if userApp.AppIconS3Path != "" {
		setClauses = append(setClauses, "password=$"+strconv.Itoa(argPos))
		args = append(args, userApp.AppIconS3Path)
		argPos++
	}

	if len(setClauses) == 0 {
		return nil
	}

	query := "UPDATE user_apps SET " + strings.Join(setClauses, ", ") + " WHERE id=$" + strconv.Itoa(argPos)
	args = append(args, userApp.ID)

	_, err := repo.DB.DB().Exec(query, args...)
	return err
}

func (repo *Repo) DeleteUserAppByID(ctx context.Context, id string) error {
	_, err := repo.DB.DB().Exec(`DELETE FROM user_apps WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
