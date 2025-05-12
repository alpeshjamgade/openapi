package service

import (
	"context"
	"open-api/internal/models"
)

func (svc *Service) CreatePlan(ctx context.Context, createPlanRequest *models.CreatePlanRequest) error {
	plan := &models.Plan{
		Name:   createPlanRequest.Name,
		Type:   createPlanRequest.Type,
		Status: createPlanRequest.Status,
		Amount: createPlanRequest.Amount,
	}

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

func (svc *Service) UpdatePlan(ctx context.Context, updatePlanRequest *models.UpdatePlanRequest) error {

	plan := &models.Plan{
		ID:     updatePlanRequest.ID,
		Name:   updatePlanRequest.Name,
		Type:   updatePlanRequest.Type,
		Amount: updatePlanRequest.Amount,
		Status: updatePlanRequest.Status,
	}

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
