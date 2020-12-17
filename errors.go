package leago

import "errors"

var (
	badRequest          = errors.New("bad request")
	unauthorized        = errors.New("unauthorized")
	forbidden           = errors.New("forbidden")
	notFound            = errors.New("data not found")
	methodNotAllowed    = errors.New("method not allowed")
	unsupportedMedia    = errors.New("unsupported media type")
	rateLimitExceeded   = errors.New("rate limit exceeded")
	internalServerError = errors.New("internal server error")
	badGateway          = errors.New("bad gateway")
	serviceUnavailable  = errors.New("service unavailable")
	gatewayTimeout      = errors.New("gateway timeout")

	responseErrors = map[int]error{
		400: badRequest,
		401: unauthorized,
		403: forbidden,
		404: notFound,
		405: methodNotAllowed,
		415: unsupportedMedia,
		429: rateLimitExceeded,
		500: internalServerError,
		502: badGateway,
		503: serviceUnavailable,
		504: gatewayTimeout,
	}

	somethingWentWrong = errors.New("something went wrong while making a request")
	tokenNotReceived   = errors.New("the token for making requests was not received")
	regionNotSupported = errors.New("riot API does not support this region")
)
