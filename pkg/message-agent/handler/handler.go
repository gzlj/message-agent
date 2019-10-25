package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gzlj/message-agent/pkg/common/global"
	"github.com/gzlj/message-agent/pkg/message-agent/infra"
	"github.com/gzlj/message-agent/pkg/message-agent/module"
)

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

	/*	url = global.HTTP_PREFIX + global.G_config.MessageCenter + global.TOKEN_URI

		req, _ = http.NewRequest("GET", url, nil)
		q := req.URL.Query()
		q.Add(global.QUERY_CLIENT_ID, global.G_config.ClientId)
		q.Add(global.QUERY_CLIENT_SECRET, global.G_config.ClientSecret)
		q.Add(global.QUERY_GRANT_TYPE, global.GRANT_TYPE_VALUE)
		req.URL.RawQuery = q.Encode()
		fmt.Println(req.URL.String())

		resp, err = http.DefaultClient.Do(req)
		defer resp.Body.Close()
		if(err != nil){
			goto FAIL
		}
		if resp.StatusCode != 200 {
			fmt.Println("resp.StatusCode: ", resp.StatusCode)
			err = errors.New(resp.Status)
			goto FAIL
		}
		token = module.TokenResponse{}
		body, _ = ioutil.ReadAll(resp.Body)
		fmt.Println("body: ", string(body))
		err =json.Unmarshal(body, &token)
		if err != nil {
			goto FAIL
		}
		fmt.Println("token struct: ", token)
		c.JSON(200, token)
		return;*/

	//c.String(200, QueryJobLogByid(jobId))
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

func HandleAlertManager(c *gin.Context) {
	//AlertManagerReqBody
	var (
		dto module.AlertManagerReqBody
		err error
	)
	fmt.Println("before c.ShouldBindJSON( ) dto: ", dto)
	if err = c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, "requet body is not correct.")
		return
	}
	fmt.Println("after c.ShouldBindJSON( ) dto: ", dto)
	c.JSON(200, dto)
}


func HandleChannel(c *gin.Context) {
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

func GetAllMsgTypes(c *gin.Context) {
	var (
		scope string
	)
	scope = c.Query("scope")
	if scope == "active" {
		c.JSON(200, module.BuildResponse(200,infra.G_ActiveChannels,"haha" ))
		return
	}

	channels, err := infra.GetAllMsgTypeNames()
	if err != nil {
		c.JSON(200, module.BuildResponse(500,channels,"" ))
		return
	}
	c.JSON(200, module.BuildResponse(200,channels,"" ))
}




