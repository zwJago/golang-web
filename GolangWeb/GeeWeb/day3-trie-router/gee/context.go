package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H map[string]interface{} 的别名
type H map[string]interface{}

// Context 上下文封装 http.ResponseWriter 和 *http.Request
// 能够高效构造 http 相应
type Context struct {
	// 源对象
	Writer http.ResponseWriter
	Req    *http.Request
	// 请求中的信息
	Path   string
	Method string
	Params map[string]string
	// 响应中的信息
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// Param 根据 key 查询路由的参数
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// PostForm 对 POST 请求，根据 key 返回 value
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 对 GET 请求，根据 key 返回 value
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置响应的状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 以 K-V 形式设置请求头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON 构造 JSON 格式响应的方法
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data 构造 Data 格式响应的方法
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 构造 HTML 格式响应的方法
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
