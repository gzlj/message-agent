package module


type MailApiBody struct {



	Title string `json:"title"`
	Timing bool `json:"timing"`
	Channels []string `json:"channels"`
	MsgType string `json:"msgType"`
	ApplyMsgType string `json:"applyMsgType"`
	Content string `json:"content"`
	CrowdType string `json:"crowdType"`
	Receiver []MessageReceiver `json:"receiver"`
}



type AlertManagerReqBody struct {
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	Annotations map[string]string `json:"annotations"`
	Labels map[string]string `json:"labels"`
	StartsAt string `json:"startsAt"`
}

type SendMessageApiBody struct {

	Title string `json:"title"`
	Timing bool `json:"timing"`
	Channels []string `json:"channels"`
	MsgType string `json:"msgType"`
	ApplyMsgType string `json:"applyMsgType"`
	Content string `json:"content"`
	CrowdType string `json:"crowdType"`
	Receiver []MessageReceiver `json:"receiver"`
}

type MessageDto struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

