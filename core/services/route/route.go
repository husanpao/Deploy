package route

import (
	"github.com/sohaha/zlsgo/znet"
	"net/http"
	"service_manager/core/helper"
)

type Meta struct {
	Title string   `json:"title"`
	Icon  string   `json:"icon"`
	Rank  int      `json:"rank"`
	Roles []string `json:"roles"`
}

type ChildMeta struct {
	Icon       string   `json:"icon"`
	ShowParent bool     `json:"showParent"`
	Title      string   `json:"title"`
	Roles      []string `json:"roles"`
	Auths      []string `json:"auths,omitempty"`
}

type Child struct {
	Path string    `json:"path"`
	Name string    `json:"name,omitempty"`
	Meta ChildMeta `json:"meta"`
}

type PermissionRouter struct {
	Path     string  `json:"path"`
	Meta     Meta    `json:"meta"`
	Children []Child `json:"children"`
}

type routeController struct {
	Prefix string
}

func NewRouteController() *routeController {
	return &routeController{}
}

func (t *routeController) Init(r *znet.Engine) error {
	// 这里可以手动绑定一些特殊路由或者中间件之类
	return nil
}

func (t *routeController) GETRoute(c *znet.Context) {
	c.JSON(http.StatusOK, helper.HttPResponse{
		Success: true,
		Data: []PermissionRouter{{Path: "/deploy", Meta: Meta{
			Title: "应用管理",
			Icon:  "raphael:icons",
			Rank:  0,
		}, Children: []Child{
			{
				Path: "/service/index",
				Name: "service",
				Meta: ChildMeta{
					Icon:       "raphael:list",
					Title:      "应用列表",
					Roles:      []string{"admin", "common"},
					ShowParent: true,
					Auths:      nil,
				},
			}, {
				Path: "/approval/index",
				Name: "approval",
				Meta: ChildMeta{
					Icon:       "raphael:anonymous",
					Title:      "发布审批",
					Roles:      []string{"admin"},
					Auths:      nil,
					ShowParent: true,
				},
			},
		}}},
		Msg: "",
	})

}
