package infra

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gzlj/message-agent/pkg/common/global"
	"github.com/gzlj/message-agent/pkg/message-agent/module"
	"io/ioutil"
	"net/http"
)

var (
	G_ActiveChannels []string
)

func SetGlobalActiveChannels(channels []string) (err error) {
	/*
		var (
			tokenResp *module.TokenResponse
		)
		tokenResp, err = GetResp(global.G_config.MessageCenter)
		if err != nil {
			return
		}
		if tokenResp.Code == 200 {
			G_ActiveToken = &tokenResp.Data
		}
		return
		FAIL:
		fmt.Println("Failed to get token when init.")
		os.Exit(-1)
	*/
	G_ActiveChannels = channels
	return
}


//func GetResp(server string) (token *module.TokenResponse, err error){
func GetAllChannels(server string) (channels []module.Channel, err error){
	var (
		req   *http.Request
		resp  *http.Response
		bytes []byte
		r     module.MessageCenterResponse
		cs []module.Channel
	)
	req = ConstructChannelReq(server)
	resp, err = http.DefaultClient.Do(req)
	if(err != nil){
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("resp.StatusCode: ", resp.StatusCode)
		err = errors.New(resp.Status)
		return
	}

	bytes, _ = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		return
	}
	fmt.Println("bytes str: ", string(bytes))
	//fmt.Println("bytes: ", r)
	if r.Code != 200 {
		err = errors.New(r.Message)
		return
	}
	bytes, _ = json.Marshal(r.Data)
	fmt.Println("bytes : ", string(bytes))
	json.Unmarshal(bytes, &cs)
	fmt.Println("cs : ", cs)
	channels = cs
	return
}

func GetAllChannelNames() (names []string, err error){
	var (
		channels []module.Channel
		channel module.Channel
	)
	channels, err = GetAllChannels(global.G_config.MessageCenter)
	if err != nil {
		return
	}
	for _, channel = range channels {
		names = append(names, channel.ChannelType)
	}
	return
}

func SetUsingChannels(channels []string) {
	G_ActiveChannels = channels
}


func ConstructChannelReq(server string) (req  *http.Request){
	server = global.HTTP_PREFIX + server + global.CHANNEL_URI
	req, _ = http.NewRequest("GET", server, nil)
	q := req.URL.Query()
	q.Add(global.QUERY_TOKEN, G_ActiveToken.AccessToken)
	req.URL.RawQuery = q.Encode()
	return
}