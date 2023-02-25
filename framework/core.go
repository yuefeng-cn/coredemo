package framework

import (
	"log"
	"net/http"
)

// 框架核心结构
type Core struct {
	// 路由
	router map[string]ControllerHandler
}

// 初始化框架核心结构
func NewCore() *Core {
	return &Core{router: map[string]ControllerHandler{}}
}

// 注册路由
func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.ServerHTTP")
	ctx := NewContext(request, response)

	// 写死测试
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.ServerHTTP")
	router(ctx)
}
