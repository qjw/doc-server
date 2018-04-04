package router

const (
	GIT_ROOT   = "/tmp/aa"
	GIT_ORIGIN = "https://github.com/qjw/test.git"
	SWAGGER_UI = "/tmp/swagger_ui"
	FRONTEND   = "/tmp/frontend"
	PORT       = 8888

	REDIS_HOST     = "localhost"
	REDIS_PORT     = 6379
	REDIS_DB       = 1
	REDIS_PASSWORD = ""

	CORP_ID           = ""
	CORP_AGENT_SECRET = ""
	CORP_AGENT_ID     = ""

	GITLAB_TOKEN = ""
	GITHUB_TOKEN = ""
)

type Config struct {
	LocalDir string `json:"localDir"`
	// 如果不存在，则不从远程Clone
	GitOrigin string `json:"gitOrigin"`
	SwaggetUi string `json:"swaggerUI"`
	Frontend  string `json:"frontend"`
	Port      int    `json:"port"`

	// Redis
	RedisHost     string `json:"redisHost""`
	RedisPort     int    `json:"redisPort"`
	RedisDb       int    `json:"redisDb"`
	RedisPassword string `json:"redisPassword"`

	// ----------------------企业号
	// 企业号corpid
	CorpID string `json:"corpID"`
	// 企业号App密钥
	CorpAgentSecret string `json:"corpAgentSecret"`
	// 因为同一个企业号会有多个Secret，这里用于区分
	CorpAgentId string `json:"corpAgentID"`

	// gitlab/github hook
	GitlabToken string `json:"gitlabToken"`
	GithubToken string `json:"githubToken"`
}

func InitConfig() *Config {
	return &Config{
		LocalDir:  GetEnv("LOCAL_DIR", GIT_ROOT),
		GitOrigin: GetEnv("GIT_ORIGIN", GIT_ORIGIN),
		SwaggetUi: GetEnv("SWAGGER_UI", SWAGGER_UI),
		Frontend:  GetEnv("FRONTEND", FRONTEND),

		Port: GetEnvInt("PORT", PORT),

		RedisHost:     GetEnv("REDIS_HOST", REDIS_HOST),
		RedisPort:     GetEnvInt("REDIS_PORT", REDIS_PORT),
		RedisDb:       GetEnvInt("REDIS_DB", REDIS_DB),
		RedisPassword: GetEnv("REDIS_PASSWORD", REDIS_PASSWORD),

		CorpID:          GetEnv("CORP_ID", CORP_ID),
		CorpAgentSecret: GetEnv("CORP_AGENT_SECRET", CORP_AGENT_SECRET),
		CorpAgentId:     GetEnv("CORP_AGENT_ID", CORP_AGENT_ID),

		GitlabToken: GetEnv("GITLAB_TOKEN", GITLAB_TOKEN),
		GithubToken: GetEnv("GITHUB_TOKEN", GITHUB_TOKEN),
	}
}
