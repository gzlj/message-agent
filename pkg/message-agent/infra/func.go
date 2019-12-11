package infra

import (
	"encoding/json"
	"fmt"
	"github.com/gzlj/message-agent/pkg/common/global"
	"github.com/gzlj/message-agent/pkg/message-agent/module"
	"io/ioutil"
	"log"
	"net/http"
	"errors"
	"strings"
)

func GetResp(server string) (token *module.TokenResponse, err error){
	var (
		resp *http.Response
		//url  string
		req  *http.Request
		body []byte

	)

	//url = global.HTTP_PREFIX + server + global.TOKEN_URI
	req = ConstructTokenReq(server)
	resp, err = http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if(err != nil){
		return
	}
	if resp.StatusCode != 200 {
		fmt.Println("resp.StatusCode: ", resp.StatusCode)
		err = errors.New(resp.Status)
		return
	}
	token = &module.TokenResponse{}
	body, _ = ioutil.ReadAll(resp.Body)
	//fmt.Println("body: ", string(body))
	err =json.Unmarshal(body, token)
	return
}

func ConstructTokenReq(server string) (req  *http.Request){
	server = global.HTTP_PREFIX + server + global.TOKEN_URI
	req, _ = http.NewRequest("GET", server, nil)
	q := req.URL.Query()
	q.Add(global.QUERY_CLIENT_ID, global.G_config.ClientId)
	q.Add(global.QUERY_CLIENT_SECRET, global.G_config.ClientSecret)
	q.Add(global.QUERY_GRANT_TYPE, global.GRANT_TYPE_VALUE)
	q.Add(global.PARAM_REDIRECT_URI, "")
	q.Add(global.PARAM_CODE, "")
	req.URL.RawQuery = q.Encode()
	return
}

func ConstructMessageReq(server string) (req  *http.Request){

	server = global.HTTP_PREFIX + server + global.MESSAGE_URI
	bodyStr := getTestMailApiBodyStr()
	req, _ = http.NewRequest("POST", server, strings.NewReader(bodyStr))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	q := req.URL.Query()
	q.Add(global.QUERY_TOKEN, G_ActiveToken.AccessToken)
	req.URL.RawQuery = q.Encode()
	return
}

func ConstructMailMessageReq(server, title, content string) (req  *http.Request){

	server = global.HTTP_PREFIX + server + global.MESSAGE_URI
	bodyStr := getMailApiBodyStr(title, content)
	req, _ = http.NewRequest("POST", server, strings.NewReader(bodyStr))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	q := req.URL.Query()
	q.Add(global.QUERY_TOKEN, G_ActiveToken.AccessToken)
	req.URL.RawQuery = q.Encode()
	return
}

func getTestMailApiBodyStr() (body string) {
	/*
	{
  "title": "告警测试test",
  "timing": true,
  "channels": [
    "YX"
  ],

  "msgType": "T",
  "applyMsgType": "GJ",
  "content": "测试数据，测试请求",
  "crowdType": "O",
  "receiver": [{
    "id": "liujun02@ly-sky.com",
    "name": "liujun02@ly-sky.com",
    "closable": true
  }]
}
	 */
	 var (
	 	b module.MailApiBody
		 bytes []byte
	 	err error
	 )
	b = module.MailApiBody{
		Title: "golang告警测试test",
		Timing: true,
		Channels: []string{"YX"},
		MsgType: "T",
		ApplyMsgType: "GJ",
		Content: "测试数据，测试请求",
		CrowdType: "O",
		Receiver: []module.MessageReceiver{module.MessageReceiver{
			Id: global.G_config.Mail,
			Name: global.G_config.Mail,
			Closable: true,
		}},
	}
	bytes, err = json.Marshal(b)
	fmt.Println("post api body: ", string(bytes))
	if err != nil {
		return ""
	}
	return string(bytes)
}

func getMailApiBodyStr(title, content string) (body string) {
	/*
	{
  "title": "告警测试test",
  "timing": true,
  "channels": [
    "YX"
  ],

  "msgType": "T",
  "applyMsgType": "GJ",
  "content": "测试数据，测试请求",
  "crowdType": "O",
  "receiver": [{
    "id": "liujun02@ly-sky.com",
    "name": "liujun02@ly-sky.com",
    "closable": true
  }]
}
	 */
	var (
		b module.MailApiBody
		bytes []byte
		err error
	)
	b = module.MailApiBody{
		Title: title,
		Timing: true,
		Channels: []string{"YX"},
		MsgType: "T",
		ApplyMsgType: "GJ",
		Content: content,
		CrowdType: "O",
		Receiver: []module.MessageReceiver{module.MessageReceiver{
			Id: global.G_config.Mail,
			Name: global.G_config.Mail,
			Closable: true,
		}},
	}
	bytes, err = json.Marshal(b)
	fmt.Println("post api body: ", string(bytes))
	if err != nil {
		return ""
	}
	return string(bytes)
}

func Mail(server, mailAddress, title, content string) (err error){
	var (
		resp *http.Response
		req  *http.Request
		body []byte
	)

	req = ConstructMessageReq(server)
	fmt.Println("req: ", req)
	resp, err = http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if(err != nil){
		fmt.Println("error when post to message center: ", err)
		return
	}

	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println("mail resp.Body: ", string(body))

	if resp.StatusCode != 200 {
		fmt.Println("resp.StatusCode: ", resp.StatusCode)
		err = errors.New(resp.Status)
	}
	return

}

func SendMessage(title, content string) (err error){
	var (
		resp *http.Response
		req  *http.Request
		//body []byte
	)
	fmt.Println("message center: ", global.G_config.MessageCenter)
	req, err = BuildSendMessageReq(global.G_config.MessageCenter, title, content)
	if err != nil {
		log.Println("Build SendMessage Request failed: "+err.Error())
		return
	}
	fmt.Println("req: ", req)
	resp, err = http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if(err != nil){
		fmt.Println("error when post to message center: ", err)
		return
	}

	//body, _ = ioutil.ReadAll(resp.Body)
	//fmt.Println("mail resp.Body: ", string(body))

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode: ", resp.StatusCode)
		err = errors.New(resp.Status)
	}
	return

}

func BuildSendMessageReq(server, title, content string) (req  *http.Request, err error){
	if G_ActiveToken == nil {
		return nil, errors.New("Token is not alived.")
	}
	server = global.HTTP_PREFIX + server + global.MESSAGE_URI
	bodyStr := getSendMessageApiBodyStr(title, content)
	req, _ = http.NewRequest("POST", server, strings.NewReader(bodyStr))
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	q := req.URL.Query()
	q.Add(global.QUERY_TOKEN, G_ActiveToken.AccessToken)
	req.URL.RawQuery = q.Encode()
	return
}

func getSendMessageApiBodyStr(title, content string) (body string) {
	/*
	{
    "title": "golang告警测试test",
    "timing": true,
    "channels": ["DX", "YX"],
    "msgType": "T",
    "applyMsgType": "GJ",
    "content": "测试数据，测试请求",
    "crowdType": "C",
    "receiver": [{
        "id": "handsome007",
        "name": "handsome",
        "closable": true
    }]
}
	 */
	var (
		b module.SendMessageApiBody
		bytes []byte
		err error
	)

	b = module.SendMessageApiBody{
		Title: title,
		Timing: true,
		Channels: G_ActiveChannels,
		MsgType: "T",
		ApplyMsgType: G_ActivMsgType,
		Content: content,
		CrowdType: "C",
		Receiver: G_Receivers,
	}
	bytes, err = json.Marshal(b)
	fmt.Println("send message api body: ", string(bytes))
	if err != nil {
		return ""
	}
	return string(bytes)
}

