package middlewares

import (
	"context"
	"net/http"
	"open-api-client/internal/constants"
	"open-api-client/internal/logger"
	"open-api-client/internal/utils"
	"time"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func (cw *CustomResponseWriter) WriteHeader(status int) {
	if cw.status == 0 {
		cw.status = status
	}

	cw.ResponseWriter.WriteHeader(status)
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := context.WithValue(r.Context(), constants.TraceID, utils.GetUUID())
		Logger := logger.CreateLoggerWithCtx(ctx)

		Logger.Infof("%s %s", r.Method, r.RequestURI)

		cw := &CustomResponseWriter{ResponseWriter: w}
		next.ServeHTTP(cw, r)

		Logger.Infof("Sent %v in %s", cw.status, time.Since(start))
		Logger.Debugf("%s", cw)
	})
}
