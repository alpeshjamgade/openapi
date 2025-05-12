package repo

import (
	"context"
	"net/http"
	"open-api/internal/client/cache"
	"open-api/internal/client/db"
	"open-api/internal/models"
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
}

type Repo struct {
	DB         db.DB
	Cache      *cache.RedisCache
	HttpClient *http.Client
}

func NewRepo(db db.DB, cache *cache.RedisCache, httpClient *http.Client) *Repo {
	return &Repo{DB: db, Cache: cache, HttpClient: httpClient}
}
