package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerResponse struct {
	Code    int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ServerResponseSuccess(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ServerResponse{
		Code:    code,
		Message: msg[code],
		Data:    data,
	})
}

func ServerResponseError(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusInternalServerError, ServerResponse{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
