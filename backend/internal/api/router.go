package api

import (
	"gotracer/internal/api/controller"
	"gotracer/internal/api/middleware"
	"gotracer/internal/ws"
	"gotracer/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)


func NewRouter() http.Handler {
	go ws.DefaultHub.Run()
	router := gin.Default()

	router.Use(middleware.CORSMiddleware())

	pG := router.Group("/api/v1/network")
	rw := response.New()

	ctl := controller.NewHandler(rw)
	ctl.SetupRoute(pG)
		
	router.GET("/ws", func(ctx *gin.Context) {
		ws.DefaultHub.ServeWS(ctx.Writer, ctx.Request)
	})

	router.Static("/assets", "../../build/web/assets")
	router.NoRoute(func(c *gin.Context) {
		c.File("../../build/web/index.html")
	})
	

	return router
}

