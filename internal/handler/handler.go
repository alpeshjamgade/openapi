package handler

import (
	"github.com/gorilla/mux"
	"net/http"
	"open-api/internal/middlewares"
	"open-api/internal/service"
)

type Handler struct {
	Service service.IService
}

func NewHandler(service service.IService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) SetupRoutes(r *mux.Router) {
	r.Use(middlewares.RequestLogger)

	// user
	r.HandleFunc("/api/v1/user", h.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/user", h.GetUserByEmail).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/user", h.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/user", h.DeleteUserByEmail).Methods(http.MethodDelete)

	// admin
	r.HandleFunc("/api/v1/admin/user", h.CreateAdmin).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/admin/user", h.GetAdminByEmail).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/admin/user", h.UpdateAdmin).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/admin/user", h.DeleteAdminByEmail).Methods(http.MethodDelete)

	// user_app
	r.HandleFunc("/api/v1/user/app", h.CreateUserApp).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/user/app", h.GetUserAppByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/user/app", h.UpdateUserApp).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/user/app", h.DeleteUserAppByID).Methods(http.MethodDelete)

	// plan
	r.HandleFunc("/api/v1/admin/plan", h.CreatePlan).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/admin/plan", h.GetPlanByID).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/admin/plan", h.UpdatePlan).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/admin/plan", h.DeletePlanByID).Methods(http.MethodDelete)

}
