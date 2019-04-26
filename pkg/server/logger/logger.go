package logger

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// LoggingResponseWriter will encapsulate a standard ResponseWritter with a copy of its statusCode
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// ResponseWriterWrapper is supposed to capture statusCode from ResponseWriter
func ResponseWriterWrapper(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

// WriteHeader is a surcharge of the ResponseWriter method
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Logger is a gorilla/mux middleware to add log to the API
func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := ResponseWriterWrapper(w)
		inner.ServeHTTP(wrapper, r)

		// if ip, _, _ := net.SplitHostPort(r.RemoteAddr); ip != os.Getenv("POD_IP") {
		if !strings.Contains(r.Header["User-Agent"][0], "kube-probe") {
			log.Printf("%s %s %s [%v] \"%s %s %s\" %d %d \"%s\" %s",
				r.RemoteAddr,
				"-",
				"-",
				start,
				r.Method,
				r.RequestURI,
				r.Proto, // string "HTTP/1.1"
				wrapper.statusCode,
				r.ContentLength,
				r.Header["User-Agent"],
				time.Since(start),
			)
		}
	})
}
