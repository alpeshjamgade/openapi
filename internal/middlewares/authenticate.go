package middlewares

import (
	"context"
	"net/http"
	"openapi-client/internal/constants"
	"openapi-client/internal/models"
	"openapi-client/internal/utils"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

		res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

		accessToken := r.Header.Get("x-access-token")
		if accessToken == "" {
			res.Status = "error"
			res.Message = "Request unauthorized"
			utils.WriteJSON(w, http.StatusUnauthorized, res)
			return
		}

		tokenSession, err := utils.ValidateAuthToken(ctx, accessToken, constants.Client)
		if err != nil {
			res.Status = "error"
			res.Message = "Request unauthorized"
			utils.WriteJSON(w, http.StatusUnauthorized, res)
			return
		}

		session, err := constants.CookieStore.Get(r, "openapi")
		if err != nil || tokenSession.UserID != session.Values["user_id"] {
			res.Status = "error"
			res.Message = "Request unauthorized"
			utils.WriteJSON(w, http.StatusUnauthorized, res)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func AuthenticateAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

		res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

		accessToken := r.Header.Get("x-access-token")
		if accessToken == "" {
			res.Status = "error"
			res.Message = "Request unauthorized"
			utils.WriteJSON(w, http.StatusUnauthorized, res)
			return
		}

		_, err := utils.ValidateAuthToken(ctx, accessToken, constants.Admin)
		if err != nil {
			res.Status = "error"
			res.Message = "Request unauthorized"
			utils.WriteJSON(w, http.StatusUnauthorized, res)
			return
		}

		//session, err := constants.CookieStore.Get(r, "openapi")
		//if err != nil || tokenSession.UserID != session.Values["user_id"] {
		//	res.Status = "error"
		//	res.Message = "Request unauthorized"
		//	utils.WriteJSON(w, http.StatusUnauthorized, res)
		//	return
		//}

		next.ServeHTTP(w, r)
	})
}
