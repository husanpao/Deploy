package core

import (
	"github.com/sohaha/zlsgo/znet"
	"service_manager/core/services/approval"
	"service_manager/core/services/login"
	"service_manager/core/services/route"
	"service_manager/core/services/service"
)

type Deploy struct {
}

func NewDeploy() *Deploy {
	return &Deploy{}
}
func (d *Deploy) Run() {
	s := _NewService()
	s.Start()
}

type Service struct {
	*znet.Engine
}

func _NewService() *Service {
	return &Service{Engine: znet.New()}
}

func (s *Service) Start() {
	s.SetMode(znet.DebugMode)
	// 异常处理
	s.Use(znet.Recovery(func(c *znet.Context, err error) {
		e := err.Error()
		c.String(500, e)
	}))
	_login := login.NewLoginController()
	s.BindStruct("", _login)
	_route := route.NewRouteController()
	s.BindStruct("", _route)
	_service := service.NewServiceController()
	s.BindStruct(_service.Prefix, _service)
	_approval := approval.NewApprovalController()
	s.BindStruct(_approval.Prefix, _approval)
	s.SetAddr(":3000")
	znet.Run()
}
