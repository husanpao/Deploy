package approval

import (
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/zdb"
)

type approvalController struct {
	Prefix string
	db     *zdb.DB
}

func NewApprovalController(db *zdb.DB) *approvalController {
	return &approvalController{Prefix: "/deploy", db: db}
}

// Init 该方法在绑定路由之前执行
func (t *approvalController) Init(r *znet.Engine) error {
	// 这里可以手动绑定一些特殊路由或者中间件之类
	return nil
}

func (t *approvalController) GETApproval(c *znet.Context) {}

func (t *approvalController) POSTApproval(c *znet.Context) {}

func (t *approvalController) IDGETApproval(c *znet.Context) {}
