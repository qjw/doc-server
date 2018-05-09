package router

const (
	LOCAL_DIR  = "/tmp/data"
	GIT_ORIGIN = "https://github.com/qjw/test.git"
	SWAGGER_UI = "/tmp/swagger_ui"
	FRONTEND   = "/tmp/frontend"
	PORT       = 8888

	GITLAB_TOKEN = ""
)

type Config struct {
	LocalDir string `json:"localDir"`
	// 如果不存在，则不从远程Clone
	GitOrigin string `json:"gitOrigin"`
	SwaggetUi string `json:"swaggerUI"`
	Frontend  string `json:"frontend"`
	Port      int    `json:"port"`

	// gitlab/github hook
	GitlabToken string `json:"gitlabToken"`
}

func InitConfig() *Config {
	return &Config{
		LocalDir:  GetEnv("LOCAL_DIR", LOCAL_DIR),
		GitOrigin: GetEnv("GIT_ORIGIN", GIT_ORIGIN),
		SwaggetUi: GetEnv("SWAGGER_UI", SWAGGER_UI),
		Frontend:  GetEnv("FRONTEND", FRONTEND),

		Port: GetEnvInt("PORT", PORT),

		GitlabToken: GetEnv("GITLAB_TOKEN", GITLAB_TOKEN),
	}
}
