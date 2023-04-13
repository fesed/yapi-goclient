package model

type SearchUserResp struct {
	CommonResp
	Data []*User `json:"data"`
}

type User struct {
	UID      int    `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	AddTime  int    `json:"addTime"`
	UpTime   int    `json:"upTime"`
}

type AddThirdUserReq struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
