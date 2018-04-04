package router

import (
	"fmt"
	"net/http"

	"github.com/qjw/go-wx-sdk/corp"
	"github.com/qjw/kelly"
	"github.com/qjw/kelly/sessions"
)

const (
	STATE = "state"
)

type LoginUrl struct {
	Url string `json:"url" binding:"required,url"`
}

func initAuthRouter(corpApi *corp.CorpApi, r kelly.Router) {
	r.GET("/login", func(c *kelly.Context) {
		user := &User{
			Name: "anonymous",
		}
		// 登录授权
		sessions.Login(c, user)

		c.WriteJson(http.StatusOK, kelly.H{
			"code":    0,
			"message": "OK",
			"data":    user,
		})
	})

	r.GET("/logout",
		sessions.LoginRequired(),
		func(c *kelly.Context) {
			sessions.Logout(c)
			c.WriteJson(http.StatusOK, kelly.H{
				"code":    0,
				"message": "OK",
			})
		})

	r.GET("/current",
		sessions.LoginRequired(),
		func(c *kelly.Context) {
			user := sessions.LoggedUser(c).(*User)
			c.WriteJson(http.StatusOK, kelly.H{
				"code":    0,
				"message": "OK",
				"data":    user,
			})
		})

	r.GET("/login_url",
		kelly.BindMiddleware(func() interface{} { return &LoginUrl{} }),
		func(c *kelly.Context) {
			params := c.GetBindParameter().(*LoginUrl)
			c.Redirect(http.StatusFound, corpApi.QrConnectUrl(
				params.Url,
				STATE,
				corpApi.Context.Config.CorpID,
				corpApi.Context.Config.Tag,
			))
		})

	r.GET("/login_qy", func(c *kelly.Context) {
		if state, err := c.GetQueryVarible("state"); err != nil || state != state {
			c.Abort(http.StatusForbidden, "")
			return
		}

		code, err := c.GetQueryVarible("code")
		if err != nil || code == "" {
			c.Abort(http.StatusForbidden, "")
			return
		}

		info, err := corpApi.Oauth2GetUserInfo(code)
		if err != nil || info == nil || !info.IsOK() {
			c.Abort(http.StatusForbidden, "")
			return
		}
		fmt.Printf("get user id %s\n", info.UserId)

		detail, err := corpApi.GetUser(info.UserId)
		if err != nil || detail == nil || !detail.IsOK() {
			c.Abort(http.StatusForbidden, "")
			return
		}

		avatar := ""
		if detail.Avatar != nil {
			avatar = *detail.Avatar
		}
		user := &User{
			Name:   detail.Name,
			Avatar: avatar,
		}
		// 登录授权
		sessions.Login(c, user)

		c.WriteJson(http.StatusOK, kelly.H{
			"code":    0,
			"message": "OK",
			"data":    user,
		})
	})
}
