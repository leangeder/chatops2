package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics is middleware to show the prometheus metrics
func Metrics(w http.ResponseWriter, r *http.Request) {
	recordMetrics()
	promhttp.Handler().ServeHTTP(w, r)
}
