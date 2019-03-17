package transport

import (
	"net/http"

	e "github.com/fzlrhmn/simple-jwt/util/error"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&ErrorResponse{
		Error: e.NotFound().Errors,
		Meta:  decorateMeta(r.Header.Get("X-Request-Id")),
	})
}
