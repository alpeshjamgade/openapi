package service

import (
	"context"
	"open-api/internal/models"
)

func (svc *Service) CreatePlan(ctx context.Context, plan *models.Plan) error {
	err := svc.repo.CreatePlan(ctx, plan)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) GetPlanByID(ctx context.Context, id string) (models.Plan, error) {
	var plan models.Plan
	plan, err := svc.repo.GetPlanByID(ctx, id)
	if err != nil {
		return plan, err
	}

	return plan, nil
}

func (svc *Service) UpdatePlan(ctx context.Context, plan *models.Plan) error {
	err := svc.repo.UpdatePlan(ctx, plan)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeletePlanByID(ctx context.Context, id string) error {
	err := svc.repo.DeletePlanByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
