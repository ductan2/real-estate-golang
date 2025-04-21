package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`    // status code
	Message string      `json:"message"` // thong bao loi
	Data    interface{} `json:"data"`    // du lai return
}

type ErrorResponseData struct {
	Code   int         `json:"code"`   // status code
	Err    string      `json:"error"`  // thong bao loi
	Detail interface{} `json:"detail"` // du lai return
}

// success response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	// message == "" set msg[code]
	if message == "" {
		message = msg[code]
	}

	// Map error codes to HTTP status codes
	var httpStatus int
	switch code {
	case BadRequest:
		httpStatus = http.StatusBadRequest
	case Unauthorized:
		httpStatus = http.StatusUnauthorized
	case Forbidden:
		httpStatus = http.StatusForbidden
	case NotFound:
		httpStatus = http.StatusNotFound
	case MethodNotAllowed:
		httpStatus = http.StatusMethodNotAllowed
	case Conflict:
		httpStatus = http.StatusConflict
	case RequestTimeout:
		httpStatus = http.StatusRequestTimeout
	case UnprocessableEntity:
		httpStatus = http.StatusUnprocessableEntity
	case TooManyRequests:
		httpStatus = http.StatusTooManyRequests
	case InternalServerError:
		httpStatus = http.StatusInternalServerError
	case NotImplemented:
		httpStatus = http.StatusNotImplemented
	case BadGateway:
		httpStatus = http.StatusBadGateway
	case ServiceUnavailable:
		httpStatus = http.StatusServiceUnavailable
	case GatewayTimeout:
		httpStatus = http.StatusGatewayTimeout
	default:
		httpStatus = http.StatusInternalServerError
	}

	c.JSON(httpStatus, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
