package api

import "github.com/gin-gonic/gin"


func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	{
		api.GET("/interfaces", listenInterfaces)
		api.POST("/scan/start", StartScan)
		api.POST("/scan/stop", StopScan)
		api.GET("/status", Status)
	}

	r.GET("/ws", ServeWS)
}
