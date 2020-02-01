package client

import "errors"

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
	httpCodes               = map[int]error{
		400: errBadRequest,
		401: errUnauthorized,
		403: errForbidden,
		404: errNotFound,
		415: errUnsupportedMediaType,
		429: errRateLimitExceeded,
		500: errInternalServerError,
		503: errServiceUnavailable,
	}
)
