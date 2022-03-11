package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

const (
	StatusCodeSuccess = 200
	StatusCodeInternalServerError = 500
	StatusCodeNotImplemented = 501
	StatusCodeBadGateway = 502
	StatusCodeServiceUnavailable = 503
	StatusCodeGatewayTimeout = 504
)

type res struct {
	Code    int64             `json:"code"`
	Message string            `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func Response(c *gin.Context, code int64, message string, data map[string]interface{}) (int, *render.JSON) {
	ctx := c.Request.Context()
	// create a done channel to tell the request it's done
	doneChan := make(chan res)

	// here you put the actual work needed for the request
	// and then send the doneChan with the status and body
	// to finish the request by writing the response
	go func() {
		result := res{
			Code:    code,
			Message: message,
			Data:    data,
		}
		if code == StatusCodeSuccess && result.Message == "" {
			result.Message = "success"
		}
		doneChan <- result
	}()

	// non-blocking select on two channels see if the request
	// times out or finishes
	select {
	// if the context is done it timed out or was cancelled
	// so don't return anything
	case <-ctx.Done():
		return StatusCodeSuccess, &render.JSON{
			Data: res{
				Code:    StatusCodeNotImplemented,
				Message: "未实现",
				Data:    data,
			},
		}
		// if the request finished then finish the request by
		// writing the response
	case res := <-doneChan:
		return StatusCodeSuccess, &render.JSON{
			Data: res,
		}
	}

}