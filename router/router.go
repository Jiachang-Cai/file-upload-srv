package router

import (
	"github.com/gin-contrib/cors"
	"net/http"

	h "fmtsmsapi/handlers"
	"fmtsmsapi/router/middlewares"
	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {
	g.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Origin", "Content-Length"},
		ExposeHeaders:    []string{"Content-Disposition"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))
	g.Use(middlewares.Logging())
	g.Use(gin.Recovery())
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	callback := g.Group("/callback")
	{
		callback.POST("/status", h.CallBackStatus) // 彩信发送回调
		callback.POST("/tmpl", h.CallBackTmpl)   // 彩信审核回调
		callback.GET("/test", h.CallBackTest)
	}
	account := g.Group("/account")
	{
		account.GET("balance", h.AccountBalance)
	}

	return g
}
