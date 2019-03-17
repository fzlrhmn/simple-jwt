package ctxhelper

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// ContextKey enum
type ContextKey int

const (
	// ContextKeyRequestID contains the ID that identifies a request.
	// In HTTP protocol, it comes from "X-Request-Id" header.
	ContextKeyRequestID ContextKey = iota
)

// PopulateFromHTTPRequest populates context values from HTTP Request
func PopulateFromHTTPRequest(ctx context.Context, r *http.Request) context.Context {
	ctx = SetRequestID(ctx, r.Header.Get("X-Request-Id"))
	return ctx
}

// GetRequestID gets request id from context
func GetRequestID(ctx context.Context) string {
	res, _ := ctx.Value(ContextKeyRequestID).(string)
	return res
}

// SetRequestID sets request id to context
// If val is empty, then generates and sets a random uuid instead
func SetRequestID(ctx context.Context, val string) context.Context {
	if val == "" {
		val = uuid.New().String()
	}
	return context.WithValue(ctx, ContextKeyRequestID, val)
}
