package service

import "github.com/sohaha/zlsgo/znet"

type serviceController struct {
	Prefix string
}

func NewServiceController() *serviceController {
	return &serviceController{Prefix: "/deploy"}
}

// Init 该方法在绑定路由之前执行
func (t *serviceController) Init(r *znet.Engine) error {
	// 这里可以手动绑定一些特殊路由或者中间件之类
	return nil
}

func (t *serviceController) GETService(c *znet.Context) {}

func (t *serviceController) POSTService(c *znet.Context) {}

func (t *serviceController) IDGETService(c *znet.Context) {}
