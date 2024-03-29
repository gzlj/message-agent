package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gzlj/message-agent/pkg/common/global"
	"github.com/gzlj/message-agent/pkg/message-agent/handler"
	"github.com/gzlj/message-agent/pkg/message-agent/infra"
	"github.com/gzlj/message-agent/pkg/message-agent/module"
	"log"
	"os"
	"runtime"
	"strings"
)

type APIServer struct {
	engine *gin.Engine
	port string
}

func (s *APIServer) Run() {
	s.engine.Run(":" + s.port)
}

// 初始化线程数量
func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func init() {

	initEnv()

	messageCenter := os.Getenv("MESSAGECENTER")
	//channel := os.Getenv("CHANNEL")
	clientId := os.Getenv("CLIENTID")
	clientSecret := os.Getenv("CLIENTSECRET")
	applyMsgType := os.Getenv("APPLYMSGTYPE")
	infra.SetUsingMsgType(applyMsgType)
	channelsStr := os.Getenv("CHANNELS")
	channels := strings.Split(channelsStr, ",")
	infra.SetUsingChannels(channels)
	port := os.Getenv("SERVERPORT")
	if port == ""{
		port="8080"
	}
	global.InitConfig(messageCenter, clientId, clientSecret, port)
	infra.InitGlobalActiveToken()
	log.Println("get token: ", infra.G_ActiveToken)
	var receivers []module.MessageReceiver
	receiversStr := os.Getenv("RECEIVERS")
	log.Println("receivers config from ENV: ", receiversStr)
	json.Unmarshal([]byte(receiversStr), &receivers)
	//log.Println("receivers: ", receivers)
	infra.SetGlobalReceivers(receivers)
}


func (s *APIServer) registryApi() {
	registryBasicApis(s.engine)
}

func registryBasicApis(r *gin.Engine) {
	//r.GET("/test", handler.TestGetToken)
	//r.GET("/mail", handler.TestPostMail)
	r.POST("/",handler.HandleAlertManager)
	r.GET("/token/get", handler.GetToken)
	r.GET("/token/isalive", handler.CheckTokenIsAlived)
	//r.GET("/channel",handler.GetChannelNames)
	//r.POST("/channel/active",handler.SetUsingChannels)
	//r.POST("/applyMsgType/active",handler.SetUsingMsgTyep)
	//r.GET("/applyMsgType",handler.GetMsgTypes)
	//r.GET("/receivers",handler.GetReceivers)
	//r.POST("/message",handler.SendMessage)
}



func main() {
	server := &APIServer{
		engine: gin.Default(),
		port: global.G_config.ServerPort,
	}
	server.registryApi()
	go infra.TtlLoop()
	server.Run()
}