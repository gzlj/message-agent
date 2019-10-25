package main
import (
	"github.com/gin-gonic/gin"
	"github.com/gzlj/message-agent/pkg/common/global"
	"github.com/gzlj/message-agent/pkg/message-agent/handler"
	"github.com/gzlj/message-agent/pkg/message-agent/infra"
	"os"
	"runtime"
)

type APIServer struct {
	engine *gin.Engine
}

// 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}


func (s *APIServer) registryApi() {
	registryBootstrap(s.engine)
}

func registryBootstrap(r *gin.Engine) {

	r.GET("/test", handler.TestGetToken)
	r.GET("/mail", handler.TestPostMail)
	r.POST("/",handler.HandleAlertManager)
	//HandleChannel
	r.GET("/channel",handler.HandleChannel)
	//SetUsingChannels
	r.POST("/channel/active",handler.SetUsingChannels)
	//GetAllMsgTypes
	r.GET("/applyMsgType",handler.GetAllMsgTypes)
}

func init() {

	initEnv()

	messageCenter := os.Getenv("MESSAGECENTER")
	channel := os.Getenv("CHANNEL")
	mail := os.Getenv("MAIL")
	phone := os.Getenv("PHONE")
	clientId := os.Getenv("CLIENTID")
	clientSecret := os.Getenv("CLIENTSECRET")

	global.InitConfig(messageCenter, clientId, clientSecret, channel, mail, phone)
	infra.InitGlobalActiveToken()

}

func main() {


	server := &APIServer{
		engine: gin.Default(),
	}
	server.registryApi()
	server.engine.Run(":18080")
}