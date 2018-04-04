package router

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

type GithubHookNotify struct {
	Ref    string `json:"ref"`
	Before string `json:"before"`
	After  string `json:"after"`
	Pusher struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher,omitempty"`
	Sender struct {
		Login string `json:"login"`
	} `json:"sender,omitempty"`
	Repository struct {
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		HtmlUrl  string `json:"html_url"`
	} `json:"repository,omitempty"`
	Commits []struct {
		ID        string `json:"id"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
		Url       string `json:"url"`
	} `json:"commits,omitempty"`
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

// 必须要登录的中间件检查
func githubRequired(token string) kelly.HandlerFunc {
	return func(c *kelly.Context) {
		event, err := c.GetHeader("X-GitHub-Event")
		if err != nil || len(event) < 1 {
			c.Abort(http.StatusForbidden, "")
			return
		}
		/*
			X-GitHub-Delivery: 21b801aa-3f01-11e8-843c-3dd69409aae4
			X-GitHub-Event: push
			X-Hub-Signature: sha1=3f068179be99e56a0b4ca80d1bd77b9587c0b295
		*/

		delivery, err := c.GetHeader("X-GitHub-Delivery")
		if err != nil || len(delivery) < 1 {
			c.Abort(http.StatusForbidden, "")
			return
		}

		signature, err := c.GetHeader("X-Hub-Signature")
		if err != nil || len(signature) < 1 {
			c.Abort(http.StatusForbidden, "")
			return
		}
		if sigs := strings.Split(signature, "="); len(sigs) != 2 || sigs[0] != "sha1" {
			c.Abort(http.StatusForbidden, "")
			return
		} else {
			signature = sigs[1]
		}

		body, _ := ioutil.ReadAll(c.Request().Body)
		mac := hmac.New(sha1.New, []byte(token))
		if _, err = mac.Write(body); err != nil {
			c.Abort(http.StatusForbidden, "")
			return
		}
		hash := mac.Sum(nil)
		if fmt.Sprintf("%x", hash) != signature {
			c.Abort(http.StatusForbidden, "")
			return
		}
		c.Set("X-GitHub-Event", event)
		c.Set("X-GitHub-Delivery", delivery)

		c.SetBody(body)
		c.InvokeNext()
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

	r.POST("/github/hook",
		githubRequired(config.GithubToken),
		func(c *kelly.Context) {

			event := c.MustGet("X-GitHub-Event")
			log.Print("event: ", event)
			body, _ := ioutil.ReadAll(c.Request().Body)
			log.Print(string(body))

			var msg GithubHookNotify
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
