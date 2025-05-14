package constants

import (
	"github.com/gorilla/sessions"
	"openapi-client/config"
)

const (
	TraceID     = "trace_id"
	Service     = "service"
	ServiceName = "openapi-client"
	Empty       = ""
	SessionKey  = "session"
	Admin       = "admin"
	Client      = "client"
)

var (
	CookieStore = sessions.NewCookieStore([]byte(config.SessionKey))
)
