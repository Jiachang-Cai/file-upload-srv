package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	h "file-upload-srv/handlers"
	"file-upload-srv/router/middlewares"
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
	callback := g.Group("/upload")
	{
		callback.POST("/add", h.UploadAdd) // 上传
	}

	return g
}
