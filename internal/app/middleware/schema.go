package middleware

import (
	"io"
	"net/http"
	"sync"
)

type contextKey string

const UserIDKey contextKey = "userID"

type CustomResponseWriter struct {
	http.ResponseWriter
	Size   int
	Status int
	once   sync.Once
	Writer io.Writer
}

func createCustomResponseWriter(w http.ResponseWriter) *CustomResponseWriter {

	return &CustomResponseWriter{ResponseWriter: w, Size: 0, Status: 200}
}

func (lrw *CustomResponseWriter) WriteHeader(code int) {
	lrw.Status = code
	lrw.once.Do(func() { lrw.ResponseWriter.WriteHeader(code) })
}

func (lrw *CustomResponseWriter) Write(body []byte) (int, error) {
	n, err := lrw.ResponseWriter.Write(body)
	lrw.Size += n
	return n, err
}
