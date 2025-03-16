package response

const (
	// 2xx Success
	Success   = 200
	Created   = 201
	Accepted  = 202
	NoContent = 204

	// 4xx Client Errors
	BadRequest          = 400
	Unauthorized        = 401
	Forbidden           = 403
	NotFound            = 404
	MethodNotAllowed    = 405
	Conflict            = 409
	RequestTimeout      = 408
	UnprocessableEntity = 422
	TooManyRequests     = 429

	// 5xx Server Errors
	InternalServerError = 500
	NotImplemented      = 501
	BadGateway          = 502
	ServiceUnavailable  = 503
	GatewayTimeout      = 504
)

var msg = map[int]string{
	// 2xx Success
	Success:   "Success",
	Created:   "Created",
	Accepted:  "Accepted",
	NoContent: "No Content",

	// 4xx Client Errors
	BadRequest:          "Bad Request",
	Unauthorized:        "Unauthorized",
	Forbidden:           "Forbidden",
	NotFound:            "Not Found",
	MethodNotAllowed:    "Method Not Allowed",
	Conflict:            "Conflict",
	RequestTimeout:      "Request Timeout",
	UnprocessableEntity: "Unprocessable Entity",
	TooManyRequests:     "Too Many Requests",

	// 5xx Server Errors
	InternalServerError: "Internal Server Error",
	NotImplemented:      "Not Implemented",
	BadGateway:          "Bad Gateway",
	ServiceUnavailable:  "Service Unavailable",
	GatewayTimeout:      "Gateway Timeout",
}
