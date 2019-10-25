package module

type Token struct {
	/*
		"access_token": "ec7050c7a0ca9e437b823a7283f8d20aed1b2b54",
	        "refresh_token": "a6b2eb2358b453e8923849fa2bcbb0ca9f98a0ea",
	        "scope": "app_role",
	        "expires_in": 2592000
	*/
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
}

type TokenResponse struct {
	Data Token `json:"data"`
	/*
		"code": 200,
	    "type": 0,
	    "pageSize": 10
	*/
	Code     int `json:"code"`
	Type     int `json:"type"`
	PageSize int `json:"pageSize"`
}

type MessageCenterResponse struct {
	Data     interface{} `json:"data"`
	Code     int         `json:"code"`
	Type     int         `json:"type"`
	PageSize int         `json:"pageSize"`
	Message string `json:"message"`
}

func BuildResponse(code int,data interface{},message string ) MessageCenterResponse {

	return MessageCenterResponse{
		Code:code,
		Data: data,
		Message: message,
	}

}
