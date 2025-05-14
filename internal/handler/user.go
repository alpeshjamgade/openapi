package handler

import (
	"context"
	"net/http"
	"openapi-client/internal/constants"
	"openapi-client/internal/logger"
	"openapi-client/internal/models"
	"openapi-client/internal/utils"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.User{}
	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	ok := h.CheckPermission(ctx, "create-user")
	if !ok {
		Logger.Errorw("invalid session")
		utils.WriteInvalidSession(w, res)
		return
	}

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

	err = h.Service.CreateUser(ctx, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Message = "User Created"
	utils.WriteJSON(w, http.StatusOK, res)

}

func (h *Handler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	queryParams := r.URL.Query()
	email := queryParams.Get("email")

	user, err := h.Service.GetUserByEmail(ctx, email)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Status = "success"
	res.Data = user
	utils.WriteJSON(w, http.StatusOK, res)
	return

}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.User{}
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

	err = h.Service.UpdateUser(ctx, req)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Status = "success"
	res.Message = "User Updated"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) DeleteUserByEmail(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	queryParams := r.URL.Query()
	email := queryParams.Get("email")
	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	err := h.Service.DeleteUserByEmail(ctx, email)
	if err != nil {
		Logger.Errorw("error reading request", "error", err)
		res.Status = "error"
		res.Message = err.Error()
		utils.WriteJSON(w, http.StatusBadRequest, res)
		return
	}

	res.Status = "success"
	res.Message = "User Deleted"
	utils.WriteJSON(w, http.StatusOK, res)
}
