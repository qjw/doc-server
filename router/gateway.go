package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"syscall"

	"github.com/qjw/kelly"
)

type GitlabHookNotify struct {
	ObjectKind string `json:"object_kind"`

	// push event
	EventName         string `json:"event_name,omitempty"`
	Ref               string `json:"ref,omitempty"`
	CheckoutSha       string `json:"checkout_sha,omitempty"`
	TotalCommitsCount int    `json:"total_commits_count,omitempty"`
	Name              string `json:"user_name,omitempty"`
	Email             string `json:"user_email,omitempty"`
	AvatarUrl         string `json:"user_avatar,omitempty"`

	User User `json:"user"`
}

// 必须要登录的中间件检查
func gitlabRequired(token string) kelly.HandlerFunc {
	return func(c *kelly.Context) {
		h, err := c.GetHeader("X-Gitlab-Token")
		if err != nil {
			c.Abort(http.StatusForbidden, "")
		} else if h != token {
			c.Abort(http.StatusForbidden, "")
		} else {
			c.InvokeNext()
		}
	}
}

func initGatewayRouter(config *Config, r kelly.Router) {
	r.POST("/gitlab/hook",
		gitlabRequired(config.GitlabToken),
		func(c *kelly.Context) {
			h, _ := c.GetHeader("X-Gitlab-Token")
			log.Print(h)
			body, _ := ioutil.ReadAll(c.Request().Body)
			log.Print(string(body))
			var msg GitlabHookNotify
			if err := json.Unmarshal(body, &msg); err != nil {
				log.Print(err.Error())
			} else {
				data, _ := json.MarshalIndent(&msg, "", " ")
				log.Print(string(data))
			}
			syscall.Kill(syscall.Getpid(), syscall.SIGUSR2)
			c.WriteString(http.StatusOK, "ok")
		})
}
