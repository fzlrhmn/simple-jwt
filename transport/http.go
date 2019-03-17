package transport

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/time/rate"

	"github.com/fzlrhmn/simple-jwt/endpoint"
	"github.com/fzlrhmn/simple-jwt/util/ctxhelper"
	e "github.com/fzlrhmn/simple-jwt/util/error"
	kitendpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	transport "github.com/go-kit/kit/transport/http"
	validator "github.com/go-playground/validator"
	"github.com/go-zoo/bone"
	"github.com/google/uuid"
	"github.com/hooqtv/glogger/adapter/newrelic"
	jsoniter "github.com/json-iterator/go"
)

var (
	json     = jsoniter.ConfigCompatibleWithStandardLibrary
	validate *validator.Validate
)

type (
	// SuccessResponse for decorate success response
	SuccessResponse struct {
		Data interface{} `json:"data"`
		Meta meta        `json:"meta"`
	}
)

// MakeHTTPHandler returns an HTTP handler for gokit endpoints
func MakeHTTPHandler(endpoints endpoint.Set) http.Handler {
	r := bone.New()
	limit := rate.NewLimiter(rate.Every(time.Minute), 1)

	serverRequestOpts := []transport.RequestFunc{
		ctxhelper.PopulateFromHTTPRequest,
	}

	serverResponseOpts := []transport.ServerResponseFunc{
		transport.SetContentType("application/vnd.api+json"),
	}

	serverOpts := []transport.ServerOption{
		transport.ServerBefore(serverRequestOpts...),
		transport.ServerAfter(serverResponseOpts...),
		transport.ServerErrorEncoder(encodeError),
	}

	r.NotFound(http.HandlerFunc(notFound))

	r.Get("/1.0/health", transport.NewServer(
		endpoints.GetHealthCheckEndpoint,
		decodeHealthCheckRequest,
		encodeResponse,
		serverOpts...,
	))

	postCreateUserEndpoint := kitendpoint.Chain(
		ratelimit.NewErroringLimiter(limit),
	)(endpoints.CreateUserEndpoint)
	r.Post("/1.0/user", transport.NewServer(
		postCreateUserEndpoint,
		decodeCreateUserRequest,
		encodeCreateResponse,
		serverOpts...,
	))

	postSigninUserEndpoint := kitendpoint.Chain(
		ratelimit.NewErroringLimiter(limit),
	)(endpoints.SigninUserEndpoint)
	r.Post("/1.0/user/signin", transport.NewServer(
		postSigninUserEndpoint,
		decodeSigninUserRequest,
		encodeSigninUserResponse,
		serverOpts...,
	))

	return r
}

// Response Encoder
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	reqID := ctxhelper.GetRequestID(ctx)

	txn := newrelic.GetTransactionFromContext(ctx)
	if txn != nil {
		txn.AddAttribute("x-request-id", reqID)
	}

	return json.NewEncoder(w).Encode(&SuccessResponse{
		Data: response,
		Meta: decorateMeta(reqID),
	})
}

func decorateMeta(reqID string) meta {
	if reqID == "" {
		reqID = uuid.New().String()
	}

	return meta{
		"now":       time.Now().UnixNano(),
		"requestId": reqID,
	}
}

func decodeAndValidate(r *http.Request, model interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		return e.RequestBodyUnparseable()
	}
	defer r.Body.Close()

	validate = validator.New()
	if err := validate.Struct(model); err != nil {
		return e.RequestBodyFailsValidation(err)
	}

	return nil
}
