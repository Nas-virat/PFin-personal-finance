package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
}

func NewErrorResponse(c *gin.Context, status int, err error) {

	c.JSON(status, Response{
		Success: false,
		Message: err.Error(),
	})
}

func NewSuccessResponse(c *gin.Context, message string, status int, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}
