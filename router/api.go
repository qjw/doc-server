package router

import (
	"net/http"

	"github.com/qjw/kelly"
	md "github.com/qjw/kelly/middleware"
	"gopkg.in/src-d/go-git.v4"
)

func Init(repository *git.Repository, r kelly.Router, config *Config) {
	// 方便前端调试，开启cors
	r = r.Group(
		"/",
		md.Cors(&md.CorsConfig{
			AllowAllOrigins: true,
			AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type"},
		}),
	)

	r.GET("/", func(c *kelly.Context) {
		c.Redirect(http.StatusFound, "/ui")
	})

	r.GET("/ui/*path",
		md.Gzip(md.BestSpeed, md.GzipMethod),
		kelly.Static(&kelly.StaticConfig{
			Dir:        http.Dir(config.Frontend),
			Indexfiles: []string{"index.html"},
		}))

	r.GET("/swagger/*path",
		md.Gzip(md.BestSpeed, md.GzipMethod),
		kelly.Static(&kelly.StaticConfig{
			Dir:        http.Dir(config.SwaggetUi),
			Indexfiles: []string{"index.html"},
		}))

	// 绑定所有的options请求来支持中间件作跨域处理
	r.OPTIONS("/*path", func(c *kelly.Context) {
		c.WriteString(http.StatusOK, "ok")
	})

	apiRouter := r.Group(
		"/api/v1",
		md.NoCache(),
	)
	initGitRouter(repository, apiRouter)
	initGatewayRouter(config, r.Group("/gateway"))
}
