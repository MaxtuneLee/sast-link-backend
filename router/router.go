package router

import (
	"github.com/NJUPT-SAST/sast-link-backend/middleware"
	"net/http"

	v1 "github.com/NJUPT-SAST/sast-link-backend/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	// FIXME: need discuss on web log
	// r.Use(middleware.WebLogger)
	r.Use(middleware.JWT)
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
		println("router:" + c.FullPath())
	})

	apiV1 := r.Group("/api/v1")
	apiV1.Use()

	usergroup := apiV1.Group("/user")
	{
		usergroup.POST("/register", v1.Register)
	}

	// admingroup := apiV1.Group("/admin")
	// {
	// }

	return r
}
