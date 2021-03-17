package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.Use(gee.Logger())
	r.Use(gee.Recovery())
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello zwjason\n")
	})

	// 数组越界
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"zwjason", "benny", "bankarian"}
		c.String(http.StatusOK, names[9])
	})

	r.Run(":8888")
}
