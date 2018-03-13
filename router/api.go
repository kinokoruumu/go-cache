package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/kinokoruumu/go-cache/controller"
	"github.com/kinokoruumu/go-cache/middleware"
)

func apiRouter(api *gin.RouterGroup) {
	api.POST("/signup", User.Create)
	api.POST("/signin", middleware.Login)

	api.GET("/articles", Article.FindAll)
}

func authApiRouter(auth *gin.RouterGroup) {
	auth.GET("/hoge", func(c *gin.Context) {
		c.JSON(200, "ok")
	})
}
