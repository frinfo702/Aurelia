package middleware

import (
	"log"
	"net/http"
)

// logging struct
type resLoggingWriter struct {
	http.ResponseWriter
	code int // error code
}

// constructor
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// override Writeheader to store actual status code
func (lw *resLoggingWriter) WriteHeader(code int) {
	lw.code = code
	lw.ResponseWriter.WriteHeader(code)
}

// actual logging processing
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request %s %s", req.Method, req.URL.Path)
		wrappedWriter := NewResLoggingWriter(w) // type casting by wrapping
		next.ServeHTTP(wrappedWriter, req)
		log.Printf("Response Status %d", wrappedWriter.code)
	})
}
