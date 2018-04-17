package router

const (
	LOCAL_DIR  = "/tmp/data"
	GIT_ORIGIN = "https://github.com/qjw/test.git"
	SWAGGER_UI = "/tmp/swagger_ui"
	FRONTEND   = "/tmp/frontend"
	PORT       = 8888

	REDIS_URL = "redis://localhost:6379/1"

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
	RedisUrl string `json:"redisUrl""`

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
		LocalDir:  GetEnv("LOCAL_DIR", LOCAL_DIR),
		GitOrigin: GetEnv("GIT_ORIGIN", GIT_ORIGIN),
		SwaggetUi: GetEnv("SWAGGER_UI", SWAGGER_UI),
		Frontend:  GetEnv("FRONTEND", FRONTEND),

		Port:     GetEnvInt("PORT", PORT),
		RedisUrl: GetEnv("REDIS_URL", REDIS_URL),

		CorpID:          GetEnv("CORP_ID", CORP_ID),
		CorpAgentSecret: GetEnv("CORP_AGENT_SECRET", CORP_AGENT_SECRET),
		CorpAgentId:     GetEnv("CORP_AGENT_ID", CORP_AGENT_ID),

		GitlabToken: GetEnv("GITLAB_TOKEN", GITLAB_TOKEN),
		GithubToken: GetEnv("GITHUB_TOKEN", GITHUB_TOKEN),
	}
}
