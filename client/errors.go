package client

import "errors"

import "fmt"

var (
	errUnimplemented        = errors.New("Unimplemented Method")
	errBadRequest           = errors.New("Bad Request")
	errUnauthorized         = errors.New("Unauthorized")
	errForbidden            = errors.New("Forbidden")
	errNotFound             = errors.New("Not Found")
	errUnsupportedMediaType = errors.New("Unsupported Media Type")
	errRateLimitExceeded    = errors.New("Rate Limit Exceeded")
	errInternalServerError  = errors.New("Internal Server Error")
	errServiceUnavailable   = errors.New("Service Unavailable")
	errBadGateway           = errors.New("Bad Gateway")
	errGatewayTimeout       = errors.New("Gateway Timeout")
	errMethodNotAllowed     = errors.New("Method Not Allowed")
	httpCodes               = map[int]error{
		400: errBadRequest,
		401: errUnauthorized,
		403: errForbidden,
		404: errNotFound,
		405: errMethodNotAllowed,
		415: errUnsupportedMediaType,
		429: errRateLimitExceeded,
		500: errInternalServerError,
		502: errBadGateway,
		503: errServiceUnavailable,
		504: errGatewayTimeout,
	}
)

//riotError is the response body riot returns with an error
type riotError struct {
	Status *riotErrorStatus `json:"status"`
}

type riotErrorStatus struct {
	Code    int32  `json:"status_code"`
	Message string `json:"message"`
}

func (r *riotError) Format() error {
	return fmt.Errorf("Error %d: %s", r.Status.Code, r.Status.Message)
}
