package login

import (
	"github.com/sohaha/zlsgo/znet"
	"net/http"
	"service_manager/core/helper"
)

type User struct {
	Username     string   `json:"username"`
	Roles        []string `json:"roles"`
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	Expires      string   `json:"expires"`
}
type TokenData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Expires      string `json:"expires"`
}
type loginController struct {
	Prefix string
}

func NewLoginController() *loginController {
	return &loginController{}
}

func (t *loginController) Init(r *znet.Engine) error {
	// 这里可以手动绑定一些特殊路由或者中间件之类
	return nil
}

func (t *loginController) POSTLogin(c *znet.Context) {
	c.JSON(http.StatusOK, helper.HttPResponse{
		Success: true,
		Data: User{
			Username:     "admin",
			Roles:        []string{"admin"},
			AccessToken:  "eyJhbGciOiJIUzUxMiJ9.admin",
			RefreshToken: "eyJhbGciOiJIUzUxMiJ9.adminRefresh",
			Expires:      "2025-01-01 00:00:00",
		},
		Msg: "",
	})

}
