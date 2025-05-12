package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"open-api/internal/constants"
	"open-api/internal/models"
	"open-api/internal/utils"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO : complete this logic
		ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())

		res := &models.HTTPResponse{Data: map[string]any{}, Status: "success", Message: ""}

		accessToken := r.Header.Get("x-access-token")
		if accessToken == "" {
			res.Status = "error"
			res.Message = "Request unauthorized"
			utils.WriteJSON(w, http.StatusUnauthorized, res)
			return
		}

		fmt.Println(accessToken)

		_, err := utils.ValidateAuthToken(ctx, accessToken)
		if err != nil {
			res.Status = "error"
			res.Message = "Request unauthorized"
			utils.WriteJSON(w, http.StatusUnauthorized, res)
			return
		}

		next.ServeHTTP(w, r)
	})
}
