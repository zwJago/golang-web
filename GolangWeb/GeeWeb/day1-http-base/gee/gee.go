package gee

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc 定义请求处理函数
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现 ServeHTTP 接口
type Engine struct {
	router map[string]HandlerFunc
}

// New 是结构体 Engine 的构造函数
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

// GET 定义如何添加 GET 请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 定义如何添加 POST 请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run 定义如何启动 http 服务器
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUNT: %s\n", req.URL)
	}
}
