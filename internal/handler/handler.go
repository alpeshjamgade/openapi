package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"openapi-client/internal/middlewares"
	"openapi-client/internal/service"
)

type Handler struct {
	Service service.IService
}

func NewHandler(service service.IService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) SetupRoutes(r *mux.Router) {
	r.Use(middlewares.RequestLogger)

	public := r.NewRoute().Subrouter()

	user := r.NewRoute().Subrouter()
	user.Use(middlewares.Authenticate)

	admin := r.NewRoute().Subrouter()
	admin.Use(middlewares.Authenticate)

	// user
	user.HandleFunc("/api/v1/user", h.CreateUser).Methods(http.MethodPost)
	user.HandleFunc("/api/v1/user", h.GetUserByEmail).Methods(http.MethodGet)
	user.HandleFunc("/api/v1/user", h.UpdateUser).Methods(http.MethodPut)
	user.HandleFunc("/api/v1/user", h.DeleteUserByEmail).Methods(http.MethodDelete)

	// admin
	admin.HandleFunc("/api/v1/admin/user", h.CreateAdmin).Methods(http.MethodPost)
	admin.HandleFunc("/api/v1/admin/user", h.GetAdminByEmail).Methods(http.MethodGet)
	admin.HandleFunc("/api/v1/admin/user", h.UpdateAdmin).Methods(http.MethodPut)
	admin.HandleFunc("/api/v1/admin/user", h.DeleteAdminByEmail).Methods(http.MethodDelete)

	// user_app
	user.HandleFunc("/api/v1/user/app", h.CreateUserApp).Methods(http.MethodPost)
	user.HandleFunc("/api/v1/user/app", h.GetUserAppByID).Methods(http.MethodGet)
	user.HandleFunc("/api/v1/user/app", h.UpdateUserApp).Methods(http.MethodPut)
	user.HandleFunc("/api/v1/user/app", h.DeleteUserAppByID).Methods(http.MethodDelete)

	// plan
	admin.HandleFunc("/api/v1/admin/plan", h.CreatePlan).Methods(http.MethodPost)
	admin.HandleFunc("/api/v1/admin/plan", h.GetPlanByID).Methods(http.MethodGet)
	admin.HandleFunc("/api/v1/admin/plan", h.UpdatePlan).Methods(http.MethodPut)
	admin.HandleFunc("/api/v1/admin/plan", h.DeletePlanByID).Methods(http.MethodDelete)

	// auth
	public.HandleFunc("/api/v1/user/login", h.Login).Methods(http.MethodPost)
	user.HandleFunc("/api/v1/user/logout", h.Logout).Methods(http.MethodPost)
	public.HandleFunc("/api/v1/admin/login", h.AdminLogin).Methods(http.MethodPost)
	user.HandleFunc("/api/v1/admin/logout", h.AdminLogin).Methods(http.MethodPost)
}
