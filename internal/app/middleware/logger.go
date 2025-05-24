package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Logger(h http.Handler) http.Handler {
	logFn := func(writer http.ResponseWriter, request *http.Request) {
		var responseStatus, responseSize int

		start := time.Now()
		customWriter := CreateCustomResponseWriter(writer)
		uri := request.RequestURI
		method := request.Method
		h.ServeHTTP(customWriter, request)
		duration := time.Since(start)
		sugar := zap.S()
		responseStatus = customWriter.Status
		responseSize = customWriter.Size
		sugar.Infoln(
			"uri", uri,
			"method", method,
			"duration", duration,
		)
		sugar.Infoln(
			"responseStatus", responseStatus,
			"responseSize", responseSize,
		)

	}
	return http.HandlerFunc(logFn)
}
