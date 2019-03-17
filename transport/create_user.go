package transport

import (
	"context"
	"net/http"

	"github.com/fzlrhmn/simple-jwt/endpoint"
	"github.com/fzlrhmn/simple-jwt/util/ctxhelper"
	e "github.com/fzlrhmn/simple-jwt/util/error"
)

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var model endpoint.CreateUserRequest
	if err := decodeAndValidate(r, &model); err != nil {
		return nil, err
	}

	if model.Username == "" {
		return nil, e.MissingPrimaryID()
	}

	return model, nil
}

// Response Encoder for create
func encodeCreateResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	reqID := ctxhelper.GetRequestID(ctx)

	w.WriteHeader(201)
	return json.NewEncoder(w).Encode(&SuccessResponse{
		Data: response,
		Meta: decorateMeta(reqID),
	})
}
