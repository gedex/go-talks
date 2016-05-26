package middleware

import (
	"net/http"
	"strings"
)

// START OMIT
func CorsHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			h := []string{
				"Authorization",
				"Content-Type",
				"Content-Range",
				"Content-Disposition",
			}

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(h, ","))
			w.Header().Set("Allow", "HEAD,GET,POST,PUT,DELETE,OPTIONS")
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// END OMIT
