package response



import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type ApiResponseWriter struct{}

func New() *ApiResponseWriter {
	return &ApiResponseWriter{}
}

func (rw *ApiResponseWriter) Success(c *gin.Context, data any, message string) {
	c.JSON(http.StatusOK, ApiResponse[any]{Status: "success", Message: message, Data: data})
}

func (rw *ApiResponseWriter) BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, ApiResponse[any]{Status: "error", Message: message})
}

func (rw *ApiResponseWriter) Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, ApiResponse[any]{Status: "error", Message: message})
}

func (rw *ApiResponseWriter) NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, ApiResponse[any]{Status: "error", Message: message})
}

func (rw *ApiResponseWriter) InternalError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, ApiResponse[any]{Status: "error", Message: message})
}