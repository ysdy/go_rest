package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ysdy/go_rest/db"
	"github.com/ysdy/go_rest/server"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World")
	})
	db.Init()
	server.Init()
	db.Close()
}
