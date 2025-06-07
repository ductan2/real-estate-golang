package middlewares

import (
	"ecommerce/pkg/response"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
)

// ErrorHandler middleware captures errors and reports them to APM
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Report panic to APM
				e := apm.DefaultTracer.NewError(fmt.Errorf("%v", err))
				e.Context.SetHTTPRequest(c.Request)
				e.Context.SetHTTPStatusCode(500)
				e.Send()

				// Send 500 response
				response.ErrorResponse(c, response.InternalServerError, "Internal Server Error")
				return
			}
		}()

		// Process request
		c.Next()

		// Check if there were any errors
		if len(c.Errors) > 0 {
			// Get the tracer
			tracer := apm.DefaultTracer

			// Create an error for each error in the context
			for _, err := range c.Errors {
				// Report error to APM
				e := tracer.NewError(err.Err)
				e.Context.SetHTTPRequest(c.Request)
				e.Context.SetHTTPStatusCode(c.Writer.Status())
				e.Send()

				// Log the error details
			}

			// If no response was sent yet, send a 500 error
			if !c.Writer.Written() {
				response.ErrorResponse(c, response.InternalServerError, "Internal Server Error")
			}
		}
	}
}

// HandleError is a helper function to add errors to gin context
func HandleError(c *gin.Context, err error, statusCode ...int) {
	// Add error to gin context
	_ = c.Error(err)

	// Get the tracer
	tracer := apm.DefaultTracer

	// Set default status code to 500 if not provided
	code := 500
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	// Report error to APM
	e := tracer.NewError(err)
	e.Context.SetHTTPRequest(c.Request)
	e.Context.SetHTTPStatusCode(code)
	e.Send()

	// Send error response with provided status code
	response.ErrorResponse(c, code, err.Error())
}
