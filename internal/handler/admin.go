package handler

import (
	"context"
	"net/http"
	"open-api/internal/constants"
	"open-api/internal/logger"
	"open-api/internal/models"
	"open-api/internal/utils"
)

func (h *Handler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.Admin{}
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

	err = h.Service.CreateAdmin(ctx, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Message = "Admin Created"
	utils.WriteJSON(w, http.StatusOK, res)

}

func (h *Handler) GetAdminByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	queryParams := r.URL.Query()
	email := queryParams.Get("email")

	admin, err := h.Service.GetAdminByEmail(ctx, email)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Status = "success"
	res.Data = admin
	utils.WriteJSON(w, http.StatusOK, res)
	return

}

func (h *Handler) UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.Admin{}
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

	err = h.Service.UpdateAdmin(ctx, req)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Status = "success"
	res.Message = "Admin Updated"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) DeleteAdminByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.HTTPRequest{}
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

}
