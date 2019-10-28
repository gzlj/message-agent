package infra

import "github.com/gzlj/message-agent/pkg/message-agent/module"

var (
	G_Receivers []module.MessageReceiver
)

func SetGlobalReceivers(receivers []module.MessageReceiver) {
	G_Receivers = receivers
}
