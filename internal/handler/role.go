package handler

import (
	"context"
	"net/http"
	"openapi-client/internal/constants"
	"openapi-client/internal/logger"
	"openapi-client/internal/models"
	"openapi-client/internal/utils"
)

func (h *Handler) CreateRole(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.Role{}
	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	err := utils.ReadJSON(w, r, req)
	if err != nil {
		Logger.Errorw("error reading request", "error", err)
		res.Status = "error"
		res.Message = "Invalid Request"
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	errs := utils.ValidateParams(req)
	if errs != nil {
		res.Status = "error"
		res.Message = errs[0].Error()
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	err = h.Service.CreateRole(ctx, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Message = "Role Created"
	utils.WriteJSON(w, http.StatusOK, res)

}

func (h *Handler) GetRoleByName(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	queryParams := r.URL.Query()
	name := queryParams.Get("name")

	role, err := h.Service.GetRoleByName(ctx, name)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Status = "success"
	res.Data = role
	utils.WriteJSON(w, http.StatusOK, res)
	return

}

func (h *Handler) UpdateRole(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.UpdateRoleRequest{}
	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	err := utils.ReadJSON(w, r, req)
	if err != nil {
		Logger.Errorw("error reading request", "error", err)
		res.Status = "error"
		res.Message = "Invalid Request"
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	errs := utils.ValidateParams(req)
	if errs != nil {
		res.Status = "error"
		res.Message = errs[0].Error()
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	err = h.Service.UpdateRole(ctx, req)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Status = "success"
	res.Message = "Role Updated"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) DeleteRoleByName(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	err := h.Service.DeleteRoleByName(ctx, name)
	if err != nil {
		Logger.Errorw("error reading request", "error", err)
		res.Status = "error"
		res.Message = err.Error()
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	res.Status = "success"
	res.Message = "Role Deleted"
	utils.WriteJSON(w, http.StatusOK, res)
}
