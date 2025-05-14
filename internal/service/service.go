package service

import (
	"context"
	"openapi-client/internal/models"
	"openapi-client/internal/repo"
)

type IService interface {
	// user
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	Login(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUserByEmail(ctx context.Context, email string) error

	// admin
	CreateAdmin(ctx context.Context, admin *models.Admin) error
	GetAdminByEmail(ctx context.Context, email string) (models.Admin, error)
	AdminLogin(ctx context.Context, loginRequest *models.LoginRequest) (*models.LoginResponse, error)
	UpdateAdmin(ctx context.Context, admin *models.Admin) error
	DeleteAdminByEmail(ctx context.Context, email string) error

	// user_app
	CreateUserApp(ctx context.Context, userApp *models.CreateUserAppRequest) error
	GetUserAppByID(ctx context.Context, id string) (models.UserApp, error)
	UpdateUserApp(ctx context.Context, userApp *models.UpdateUserAppRequest) error
	DeleteUserAppByID(ctx context.Context, id string) error

	// plan
	CreatePlan(ctx context.Context, plan *models.CreatePlanRequest) error
	GetPlanByID(ctx context.Context, id string) (models.Plan, error)
	UpdatePlan(ctx context.Context, plan *models.UpdatePlanRequest) error
	DeletePlanByID(ctx context.Context, id string) error

	// role
	CreateRole(ctx context.Context, role *models.Role) error
	GetRoleByName(ctx context.Context, name string) (models.Role, error)
	UpdateRole(ctx context.Context, role *models.UpdateRoleRequest) error
	DeleteRoleByName(ctx context.Context, name string) error

	// permission
	CreatePermission(ctx context.Context, role *models.Permission) error
	GetPermissionByName(ctx context.Context, name string) (models.Permission, error)
	UpdatePermission(ctx context.Context, role *models.UpdatePermissionRequest) error
	DeletePermissionByName(ctx context.Context, name string) error

	// role permission
	CreateRolePermission(ctx context.Context, rolePermission *models.RolePermission) error
}

type Service struct {
	repo repo.IRepo
}

func NewService(repo repo.IRepo) *Service {
	return &Service{repo: repo}
}
