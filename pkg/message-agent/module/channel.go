package module

type Channel struct {
	/*
			{
		            "id": "IM",
		            "name": "IM",
		            "channelType": "IM"
		        },
	*/

	Id          string `json:"id"`
	Name        string `json:"name"`
	ChannelType string `json:"channelType"`
}
