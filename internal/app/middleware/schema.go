package middleware

import (
	"io"
	"net/http"
	"sync"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	Size   int
	Status int
	once   sync.Once
	Writer io.Writer
}

func CreateCustomResponseWriter(w http.ResponseWriter) *CustomResponseWriter {

	return &CustomResponseWriter{ResponseWriter: w, Size: 0, Status: 200, Writer: nil}
}

func (lrw *CustomResponseWriter) WriteHeader(code int) {
	lrw.Status = code
	lrw.once.Do(func() { lrw.ResponseWriter.WriteHeader(code) })
}

func (lrw *CustomResponseWriter) Write(body []byte) (int, error) {
	var n int
	var err error
	if lrw.Writer != nil {
		n, err = lrw.Writer.Write(body)
	} else {
		n, err = lrw.ResponseWriter.Write(body)
	}
	lrw.Size += n
	return n, err
}
