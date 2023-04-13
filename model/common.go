package model

const (
	StatusOK = 0
)

type Claims struct {
	UID int `json:"uid"`
	Iat int `json:"iat"`
	Exp int `json:"exp"`
}

type CommonResp struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type DeleteReq struct {
	Id int `json:"id"`
}

type AddResp struct {
	CommonResp
	Data struct {
		ID int `json:"_id"`
	} `json:"data"`
}
