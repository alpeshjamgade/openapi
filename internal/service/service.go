package service

import (
	"context"
	"open-api/internal/models"
	"open-api/internal/repo"
)

type IService interface {
	// user
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUserByEmail(ctx context.Context, email string) error

	// admin
	CreateAdmin(ctx context.Context, admin *models.Admin) error
	GetAdminByEmail(ctx context.Context, email string) (models.Admin, error)
	UpdateAdmin(ctx context.Context, admin *models.Admin) error
	DeleteAdminByEmail(ctx context.Context, email string) error

	// user_app
	CreateUserApp(ctx context.Context, userApp *models.UserApp) error
	GetUserAppByID(ctx context.Context, id string) (models.UserApp, error)
	UpdateUserApp(ctx context.Context, userApp *models.UserApp) error
	DeleteUserAppByID(ctx context.Context, id string) error

	// plan
	CreatePlan(ctx context.Context, plan *models.Plan) error
	GetPlanByID(ctx context.Context, id string) (models.Plan, error)
	UpdatePlan(ctx context.Context, plan *models.Plan) error
	DeletePlanByID(ctx context.Context, id string) error
}

type Service struct {
	repo repo.IRepo
}

func NewService(repo repo.IRepo) *Service {
	return &Service{repo: repo}
}
