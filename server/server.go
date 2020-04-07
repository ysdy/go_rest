package server

import (
	"github.com/gin-gonic/gin"

	"github.com/ysdy/go_rest/controller"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	p := r.Group("/pets")
	{
		ctrl := pet.Controller{}
		p.GET("", ctrl.Index)
		p.GET("/:id", ctrl.Show)
		p.POST("", ctrl.Create)
		// p.PUT("/:id", ctrl.Update)
		// p.DELETE("/:id", ctrl.Delete)
	}

	return r
}
