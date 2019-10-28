package module

type MessageReceiver struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Closable bool `json:"closable"`
}