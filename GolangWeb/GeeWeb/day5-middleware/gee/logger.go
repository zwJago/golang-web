package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// 开启计时器
		t := time.Now()
		// 处理请求
		c.Next()
		// 计算耗时
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
