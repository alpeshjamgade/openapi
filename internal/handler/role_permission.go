package handler

import (
	"context"
	"net/http"
	"openapi-client/internal/constants"
	"openapi-client/internal/logger"
	"openapi-client/internal/models"
	"openapi-client/internal/utils"
)

func (h *Handler) CreateRolePermission(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.RolePermission{}
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

	err = h.Service.CreateRolePermission(ctx, req)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Message = "Role Permission Created"
	utils.WriteJSON(w, http.StatusOK, res)

}
