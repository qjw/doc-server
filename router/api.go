package router

import (
	"fmt"
	"net/http"

	"github.com/qjw/go-wx-sdk/cache"
	"github.com/qjw/go-wx-sdk/corp"
	"github.com/qjw/kelly"
	md "github.com/qjw/kelly/middleware"
	"github.com/qjw/kelly/sessions"
	"gopkg.in/redis.v5"
	"gopkg.in/src-d/go-git.v4"
)

func initRedis(config *Config) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.RedisHost, config.RedisPort),
		Password: config.RedisPassword,
		DB:       config.RedisDb,
	})
	if err := redisClient.Ping().Err(); err != nil {
		panic("failed to connect redis")
	}
	return redisClient
}

func initStore(config *Config, redisClient *redis.Client) sessions.Store {
	store, err := sessions.NewRediStore(redisClient, []byte("abcdefg"))
	if err != nil {
		panic(err)
	}
	return store
}

func Init(repository *git.Repository, r kelly.Router, config *Config) {
	// redis客户端
	redisClient := initRedis(config)
	// session
	store := initStore(config, redisClient)
	// 企业号
	corpContext := corp.NewContext(
		&corp.Config{
			CorpID:     config.CorpID,
			CorpSecret: config.CorpAgentSecret,
			Tag:        config.CorpAgentId,
		},
		cache.NewCache(redisClient),
	)
	corpApi := corp.NewCorpApi(corpContext)

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
		sessions.SessionMiddleware(store, sessions.AUTH_SESSION_NAME),
		sessions.AuthMiddleware(&sessions.AuthOptions{
			User: &User{},
		}),
	)
	initGitRouter(repository, apiRouter)
	initAuthRouter(corpApi, apiRouter)
	initGatewayRouter(config, r.Group("/gateway"))
}
