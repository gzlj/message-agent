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

type MessageReceiver struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Closable bool `json:"closable"`
}

type AlertManagerReqBody struct {
	Alerts []Alert `json:"alerts"`
}

type Alert struct {
	Annotations map[string]string `json:"annotations"`
	Labels map[string]string `json:"labels"`
	StartsAt string `json:"startsAt"`
}