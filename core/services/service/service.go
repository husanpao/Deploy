package service

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/zdb"
)

type serviceController struct {
	Prefix string
	db     *zdb.DB
}

func NewServiceController(db *zdb.DB) *serviceController {
	return &serviceController{Prefix: "/deploy", db: db}
}

// Init 该方法在绑定路由之前执行
func (t *serviceController) Init(r *znet.Engine) error {
	// 这里可以手动绑定一些特殊路由或者中间件之类
	return nil
}

func (t *serviceController) GETService(c *znet.Context) {}

func (t *serviceController) POSTService(c *znet.Context) {}

func (t *serviceController) IDGETService(c *znet.Context) {}
