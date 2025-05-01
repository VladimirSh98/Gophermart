package middleware

import (
	"compress/gzip"
	"net/http"
	"strings"
)

func Compress(h http.Handler) http.Handler {
	logFn := func(writer http.ResponseWriter, request *http.Request) {
		var err error

		customCompressWriter := CreateCustomResponseWriter(writer)

		contentEncoding := request.Header.Get("Content-Encoding")
		sendCompress := strings.Contains(contentEncoding, "gzip")
		if sendCompress {
			var cr *gzip.Reader
			cr, err = gzip.NewReader(request.Body)
			if err != nil {
				customCompressWriter.WriteHeader(http.StatusInternalServerError)
				return
			}
			request.Body = cr
			defer cr.Close()
		}
		if strings.Contains(request.Header.Get("Accept-Encoding"), "gzip") {
			gzipWriter := gzip.NewWriter(customCompressWriter)
			defer gzipWriter.Close()
			customCompressWriter.Writer = gzipWriter
			customCompressWriter.Header().Set("Content-Encoding", "gzip")
			h.ServeHTTP(customCompressWriter, request)
		}
	}
	return http.HandlerFunc(logFn)
}
