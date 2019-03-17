package transport

import (
	"context"
	"net/http"

	"github.com/fzlrhmn/simple-jwt/util/ctxhelper"
	e "github.com/fzlrhmn/simple-jwt/util/error"
	"github.com/spf13/viper"
)

type (
	// ErrorResponse for hold error response
	ErrorResponse struct {
		Error []e.Error `json:"errors"`
		Meta  meta      `json:"meta"`
	}

	meta map[string]interface{}
)

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/vnd.api+json")

	eo, ok := err.(e.Container)
	if !ok {
		eo = e.Unexpected().Wrap(err)
	}

	statusCode := eo.GetFirstError().Status
	if statusCode == 0 {
		statusCode = 500
	}

	env := viper.GetString("app.env")
	if env == "production" {
		// If in production env, don't leak the stack trace
		eo.RemoveTraces()
	}

	reqID := ctxhelper.GetRequestID(ctx)

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&ErrorResponse{
		Error: eo.Errors,
		Meta:  decorateMeta(reqID),
	})
}
