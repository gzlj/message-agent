package infra

import (
	"github.com/gzlj/message-agent/pkg/common/global"
	"github.com/gzlj/message-agent/pkg/message-agent/module"
	"time"
)

var (
	G_ActiveToken *module.Token
)

func InitGlobalActiveToken() (err error){
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
	/*FAIL:
	fmt.Println("Failed to get token when init.")
	os.Exit(-1)
	*/
}

func TtlLoop() {

	timer1 := time.NewTimer(5 * time.Second)
	for {

		select {
		case <-timer1.C:
			checkActiveToken()
			timer1.Reset(5 * time.Second)
		}
	}

}

func checkActiveToken() {
	if G_ActiveToken != nil {
		G_ActiveToken.ExpiresIn = G_ActiveToken.ExpiresIn - 5
		if G_ActiveToken.ExpiresIn < 600 {
			InitGlobalActiveToken()
		}
		return
	}
	InitGlobalActiveToken()
}

func TokenIsAlive() bool {
	if G_ActiveToken != nil && G_ActiveToken.ExpiresIn >= 600 {
		return true
	}
	return false
}
