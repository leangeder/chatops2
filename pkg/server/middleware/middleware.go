package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// setupGlobalMiddleware will setup CORS
func SetupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}
