package global

// 程序配置
type Config struct {
	MessageCenter string `json:"messageCenter"`
	ClientId      string `json:"clientId"`
	ClientSecret  string `json:"clientSecret"`
	Channel       string `json:"channel"`
	Mail          string `json:"mail"`
	Phone         string `json:"phone"`
}

var (
	// 单例
	G_config *Config
)

var (
	HTTP_PREFIX         = "http://"
	QUERY_CLIENT_ID     = "client_id"
	QUERY_CLIENT_SECRET = "client_secret"
	QUERY_GRANT_TYPE    = "grant_type"
	QUERY_TOKEN = "token"
	GRANT_TYPE_VALUE    = "client_credentials"

	TOKEN_URI    ="/api/v4/oauth/token"
	MESSAGE_URI   = "/api/v4/oauth/message"
)

// 加载配置
func InitConfig(messageCenter, clientId, clientSecret, channel, mail, phone string) (err error) {
	conf := Config{
		MessageCenter: messageCenter,
		ClientId:      clientId,
		ClientSecret:  clientSecret,
		Channel:       channel,
		Mail:          mail,
		Phone:         phone,
	}
	// 3, 赋值单例
	G_config = &conf
	return
}
