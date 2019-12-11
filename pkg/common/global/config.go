package global

// 程序配置
type Config struct {
	MessageCenter string   `json:"messageCenter"`
	ClientId      string   `json:"clientId"`
	ClientSecret  string   `json:"clientSecret"`
	Channel       string   `json:"channel"`
	Mail          string   `json:"mail"`
	Phone         string   `json:"phone"`
	AllChannels   []string `json:"allChannels"`
	UsingChannels []string `json:"usingChannels"`
	AllMsgTypes   []string `json:"allMsgTypes"`
	ActiveMsgType []string `json:"activeMsgType"`
	ServerPort string   `json:"serverPort"`
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
	PARAM_REDIRECT_URI = "redirect_uri"
	PARAM_CODE = "code"

	TOKEN_URI    ="/api/v4/oauth/token"
	MESSAGE_URI   = "/api/v4/oauth/message"
	CHANNEL_URI = "/api/v4/oauth/channels"
	MSSAGE_TYPE_URI = "/api/v4/oauth/applyMsgTypes"


)

// 加载配置
func InitConfig(messageCenter, clientId, clientSecret, serverPort string) (err error) {
	conf := Config{
		MessageCenter: messageCenter,
		ClientId:      clientId,
		ClientSecret:  clientSecret,
		ServerPort: serverPort,
	}
	// 3, 赋值单例
	G_config = &conf
	return
}
