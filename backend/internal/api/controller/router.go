package controller

import "github.com/gin-gonic/gin"


func (h *handler) SetupRoute(router *gin.RouterGroup) {
	router.GET("/interfaces", h.getNetInterfaces)
	router.GET("/domain", h.getDomainIP)
}