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
	G_ActivMsgType string
)

func SetGlobalActivMsgTypes(msgTypes string) (err error) {

	G_ActivMsgType = msgTypes
	return
}

func GetAllMsgTypes(server string) (msgTypes []module.MessageType, err error){
	var (

		req   *http.Request
		resp  *http.Response
		bytes []byte
		r     module.MessageCenterResponse
		tmp   []module.MessageType
	)
	req = ConstructMsgTypeReq(server)
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
	json.Unmarshal(bytes, &tmp)
	fmt.Println("cs : ", tmp)
	msgTypes = tmp
	return
}

func GetAllMsgTypeNames() (names []string, err error){
	var (
		msgTypes []module.MessageType
		msgType module.MessageType
	)
	msgTypes, err = GetAllMsgTypes(global.G_config.MessageCenter)
	if err != nil {
		return
	}
	for _, msgType = range msgTypes {
		names = append(names, msgType.Code)
	}
	return
}

func ConstructMsgTypeReq(server string) (req  *http.Request){
	server = global.HTTP_PREFIX + server + global.MSSAGE_TYPE_URI
	req, _ = http.NewRequest("GET", server, nil)
	q := req.URL.Query()
	q.Add(global.QUERY_TOKEN, G_ActiveToken.AccessToken)
	req.URL.RawQuery = q.Encode()
	return
}

func SetUsingMsgType(msgType string) {

	G_ActivMsgType = msgType
}


