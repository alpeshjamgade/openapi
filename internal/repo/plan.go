package repo

import (
	"context"
	"open-api/internal/models"
	"strconv"
	"strings"
)

func (repo *Repo) CreatePlan(ctx context.Context, plan *models.Plan) error {

	_, err := repo.DB.DB().Exec(
		`INSERT INTO plans(name, type, status, amount) VALUES($1, $2, $3, $4)`,
		plan.Name,
		plan.Type,
		plan.Status,
		plan.Amount,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *Repo) GetPlanByID(ctx context.Context, id string) (models.Plan, error) {
	var plan models.Plan
	sqlRow := repo.DB.DB().QueryRow(`SELECT id, name, type, status, amount, created_at, updated_at FROM plans WHERE id = $1`, id)

	err := sqlRow.Scan(
		&plan.ID,
		&plan.Name,
		&plan.Type,
		&plan.Status,
		&plan.Amount,
		&plan.CreatedAt,
		&plan.UpdatedAt,
	)
	if err != nil {
		return plan, err
	}

	return plan, nil
}

func (repo *Repo) UpdatePlan(ctx context.Context, plan *models.Plan) error {
	setClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if plan.Name != "" {
		setClauses = append(setClauses, "name=$"+strconv.Itoa(argPos))
		args = append(args, plan.Name)
		argPos++
	}
	if plan.Type != "" {
		setClauses = append(setClauses, "type=$"+strconv.Itoa(argPos))
		args = append(args, plan.Status)
		argPos++
	}
	if plan.Status != "" {
		setClauses = append(setClauses, "status=$"+strconv.Itoa(argPos))
		args = append(args, plan.Status)
		argPos++
	}
	if plan.Amount != 0 {
		setClauses = append(setClauses, "amount=$"+strconv.Itoa(argPos))
		args = append(args, plan.Amount)
		argPos++
	}

	if len(setClauses) == 0 {
		return nil
	}

	query := "UPDATE plans SET " + strings.Join(setClauses, ", ") + " WHERE id=$" + strconv.Itoa(argPos)
	args = append(args, plan.ID)

	_, err := repo.DB.DB().Exec(query, args...)
	return err
}

func (repo *Repo) DeletePlanByID(ctx context.Context, id string) error {
	_, err := repo.DB.DB().Exec(`DELETE FROM plans WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
