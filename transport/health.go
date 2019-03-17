package transport

import (
	"context"
	"net/http"
)

// Request Decoder
func decodeHealthCheckRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
