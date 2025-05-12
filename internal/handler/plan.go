package handler

import (
	"context"
	"net/http"
	"open-api/internal/constants"
	"open-api/internal/logger"
	"open-api/internal/models"
	"open-api/internal/utils"
)

func (h *Handler) CreatePlan(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.CreatePlanRequest{}
	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	err := utils.ReadJSON(w, r, req)
	if err != nil {
		Logger.Errorf("error reading request, error: %s", err)
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

	plan := &models.Plan{
		Name:   req.Name,
		Type:   req.Type,
		Status: req.Status,
		Amount: req.Amount,
	}
	err = h.Service.CreatePlan(ctx, plan)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Message = "Plan Created"
	utils.WriteJSON(w, http.StatusOK, res)

}

func (h *Handler) GetPlanByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	plan, err := h.Service.GetPlanByID(ctx, id)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Status = "success"
	res.Data = plan
	utils.WriteJSON(w, http.StatusOK, res)
	return

}

func (h *Handler) UpdatePlan(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
	Logger := logger.CreateFileLoggerWithCtx(ctx)

	req := &models.UpdatePlanRequest{}
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

	plan := &models.Plan{
		ID:     req.ID,
		Name:   req.Name,
		Type:   req.Type,
		Amount: req.Amount,
		Status: req.Status,
	}
	err = h.Service.UpdatePlan(ctx, plan)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Status = "success"
	res.Message = "Plan Updated"
	utils.WriteJSON(w, http.StatusOK, res)
	return
}

func (h *Handler) DeletePlanByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

	res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	err := h.Service.DeletePlanByID(ctx, id)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	res.Status = "success"
	res.Message = "Plan Deleted"
	utils.WriteJSON(w, http.StatusOK, res)
	return

}
