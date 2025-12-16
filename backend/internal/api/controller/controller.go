package controller


import (
	"gotracer/pkg/response"
)
type handler struct {
	response *response.ApiResponseWriter
}

func NewHandler(resp *response.ApiResponseWriter) *handler {
	return &handler{
		response: resp,
	}
}


