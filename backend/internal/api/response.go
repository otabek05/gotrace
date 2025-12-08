package api

import (
	"gotrace/backend/internal/types"

	"github.com/gin-gonic/gin"
)

func success[T any](c *gin.Context, data T) {
	c.JSON(200, types.ApiResponse[any]{
		Success: true,
		Message: "",
		Data: data,
	})
}

func error(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, types.ApiResponse[any]{
		Success: false,
		Message: message,
	})
}
