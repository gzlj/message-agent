package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gzlj/message-agent/pkg/message-agent/infra"
	"github.com/gzlj/message-agent/pkg/message-agent/module"
)

/*
func TestGetToken(c *gin.Context) {
	var (
		err  error
		tokenResp *module.TokenResponse
	)

	tokenResp, err = infra.GetResp(global.G_config.MessageCenter)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, tokenResp)
}



func TestPostMail(c *gin.Context) {
	if ! infra.TokenIsAlive() {
		infra.InitGlobalActiveToken()
		if ! infra.TokenIsAlive() {
			c.JSON(500, "token is invalied.")
			return
		}
	}
	// do mail
	err := infra.Mail(global.G_config.MessageCenter, global.G_config.Mail,"go test", "test content")
	fmt.Println("Mail() error: ", err)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, "mail job is stared OK.")

}
*/


func HandleAlertManager(c *gin.Context) {
	//AlertManagerReqBody
	var (
		dto module.AlertManagerReqBody
		err error
		bytes []byte
	)
	fmt.Println("before c.ShouldBindJSON( ) dto: ", dto)
	if err = c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, "requet body is not correct.")
		return
	}
	bytes, err = json.Marshal(dto.Alerts)
	fmt.Println("prometheus alert: ", string(bytes))
	infra.SendMessage("prometheus alert", string(bytes))
	c.JSON(200, nil)

	//fmt.Println("after c.ShouldBindJSON( ) dto: ", dto)
	//c.JSON(200, dto)
}

/*
func GetChannelNames(c *gin.Context) {
	var (
		scope string
	)
	scope = c.Query("scope")
	if scope == "active" {
		c.JSON(200, module.BuildResponse(200,infra.G_ActiveChannels,"haha" ))
		return
	}
	channels, err := infra.GetAllChannelNames()
	if err != nil {
		c.JSON(200, module.BuildResponse(500,channels,"" ))
		return
	}
	c.JSON(200, module.BuildResponse(200,channels,"" ))
}

func SetUsingChannels(c *gin.Context) {
	var (
		channels []string
		err error
	)
	if err = c.ShouldBindJSON(&channels); err != nil {
		c.JSON(400, "requet body is not correct.")
		return
	}

	infra.SetUsingChannels(channels)

	c.JSON(200, nil)
}

func SetUsingMsgTyep(c *gin.Context) {
	var (
		msgType string
	)
	msgType = c.Query("msgType")

	if msgType == "" {
		c.JSON(400, "please specify MsgTyep.")
		return
	}
	infra.SetUsingMsgType(msgType)

	c.JSON(200, nil)
}




func GetMsgTypes(c *gin.Context) {
	var (
		scope string
	)
	scope = c.Query("scope")
	if scope == "active" {
		c.JSON(200, module.BuildResponse(200,[]string{infra.G_ActivMsgType},"haha" ))
		return
	}

	names, err := infra.GetAllMsgTypeNames()
	if err != nil {
		c.JSON(200, module.BuildResponse(500, names,"" ))
		return
	}
	c.JSON(200, module.BuildResponse(200, names,"" ))
}

func GetReceivers(c *gin.Context) {

		c.JSON(200, module.BuildResponse(200,infra.G_Receivers,"" ))
		return

}


//SendMessage
func SendMessage(c *gin.Context) {

	var (
		dto module.MessageDto
		err error
	)

	if err = c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, "requet body is not correct.")
		return
	}
	fmt.Println("dto: ", dto)
	infra.SendMessage(dto.Title, dto.Content)
	c.JSON(200, nil)
	return

}

*/


