package core

import (
	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/znet"
	"github.com/zlsgo/conf"
	"service_manager/core/config"
	"service_manager/core/helper"
	"service_manager/core/services/approval"
	"service_manager/core/services/login"
	"service_manager/core/services/route"
	"service_manager/core/services/service"
)

type Deploy struct {
	cfg []string
}

func NewDeploy(cfg ...string) *Deploy {
	return &Deploy{cfg: cfg}
}
func (d *Deploy) Run() {
	s := _NewService(d.cfg...)
	s.Start()
}

type Service struct {
	*znet.Engine
	config config.Config
}

func _NewService(cfg ...string) *Service {

	return &Service{Engine: znet.New(), config: parseConfig(cfg...)}
}

func parseConfig(cfg ...string) (c config.Config) {
	cfile := "./deploy.toml"
	if len(cfg) > 0 {
		cfile = cfg[0]
	}
	cfgObj := conf.New(cfile)
	err := cfgObj.Read()
	if err != nil {
		zlog.Panic(err)
	}
	c.DB = config.DataBaseConfig{
		Host:     cfgObj.Get("database.host").String(),
		Port:     cfgObj.Get("database.port").Int(),
		User:     cfgObj.Get("database.username").String(),
		Password: cfgObj.Get("database.password").String(),
		DBName:   cfgObj.Get("database.dbname").String(),
	}
	return
}
func (s *Service) Start() {
	mysqlContext := helper.NewMysqlContext(s.config.DB)
	s.SetMode(znet.DebugMode)
	// 异常处理
	s.Use(znet.Recovery(func(c *znet.Context, err error) {
		e := err.Error()
		c.String(500, e)
	}))
	_login := login.NewLoginController(mysqlContext)
	s.BindStruct("", _login)
	_route := route.NewRouteController(mysqlContext)
	s.BindStruct("", _route)
	_service := service.NewServiceController(mysqlContext)
	s.BindStruct(_service.Prefix, _service)
	_approval := approval.NewApprovalController(mysqlContext)
	s.BindStruct(_approval.Prefix, _approval)
	s.SetAddr(":3000")
	znet.Run()
}
