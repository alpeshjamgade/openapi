package repo

import (
	"context"
	"net/http"
	"openapi-client/internal/client/cache"
	"openapi-client/internal/client/db"
	"openapi-client/internal/models"
)

type IRepo interface {
	// user
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	GetUserByEmailAndPassword(ctx context.Context, email, password string) (models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUserByEmail(ctx context.Context, email string) error

	// admin
	CreateAdmin(ctx context.Context, admin *models.Admin) error
	GetAdminByEmail(ctx context.Context, email string) (models.Admin, error)
	GetAdminByEmailAndPassword(ctx context.Context, email, password string) (models.Admin, error)
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
	CreateRolePermission(ctx context.Context, role *models.RolePermission) error
}

type Repo struct {
	DB         db.DB
	Cache      *cache.RedisCache
	HttpClient *http.Client
}

func NewRepo(db db.DB, cache *cache.RedisCache, httpClient *http.Client) *Repo {
	return &Repo{DB: db, Cache: cache, HttpClient: httpClient}
}
