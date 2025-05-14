package utils

import (
	"context"
	"openapi-client/internal/constants"
	"openapi-client/internal/models"
)

func GetSessionFromContext(ctx context.Context) (*models.Session, bool) {
	session, ok := ctx.Value(constants.SessionKey).(*models.Session)
	return session, ok
}
