package handler

import (
	"context"
	"net/http"
	"openapi-client/internal/constants"
	"openapi-client/internal/logger"
	"openapi-client/internal/models"
	"openapi-client/internal/utils"
)

func (h *Handler) CreatePermission(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.Permission{}
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

	err = h.Service.CreatePermission(ctx, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Message = "Permission Created"
	utils.WriteJSON(w, http.StatusOK, res)

}

func (h *Handler) GetPermissionByName(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	queryParams := r.URL.Query()
	name := queryParams.Get("name")

	permission, err := h.Service.GetPermissionByName(ctx, name)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Status = "success"
	res.Data = permission
	utils.WriteJSON(w, http.StatusOK, res)
	return

}

func (h *Handler) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.UpdatePermissionRequest{}
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

	err = h.Service.UpdatePermission(ctx, req)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Status = "success"
	res.Message = "Permission Updated"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) DeletePermissionByName(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	err := h.Service.DeletePermissionByName(ctx, name)
	if err != nil {
		Logger.Errorw("error reading request", "error", err)
		res.Status = "error"
		res.Message = err.Error()
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	res.Status = "success"
	res.Message = "Permission Deleted"
	utils.WriteJSON(w, http.StatusOK, res)
}
