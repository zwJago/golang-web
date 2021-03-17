package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 定义请求的处理函数
type HandlerFunc func(*Context)

// RouterGroup 路由分组
type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

// Engine 实现 Handler 接口的结构体
type Engine struct {
	router *router
	*RouterGroup
	groups []*RouterGroup
}

// New Engine的构造函数
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 用来创建一个新的 RouterGroup
// 所有的 groups 共享一个 Engine 实例
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// GET 定义如何添加 GET 请求
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 定义如何添加 POST 请求
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

// Run 包装 ListenAndServe 方法启动 http 服务器
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
