package constants

import (
	"github.com/gorilla/sessions"
	"open-api-client/config"
)

const (
	TraceID     = "trace_id"
	Service     = "service"
	ServiceName = "open-api-client"
	Empty       = ""
)

var (
	CookieStore = sessions.NewCookieStore([]byte(config.SessionKey))
)
