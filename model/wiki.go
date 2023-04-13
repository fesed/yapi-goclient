package model

type GetWikiResp struct {
	CommonResp
	Data *Wiki `json:"data"`
}

type Wiki struct {
	EditUID   int    `json:"edit_uid"`
	ID        int    `json:"_id"`
	ProjectID int    `json:"project_id"`
	Desc      string `json:"desc"`
	Markdown  string `json:"markdown"`
	Username  string `json:"username"`
	UID       int    `json:"uid"`
	AddTime   int    `json:"add_time"`
	UpTime    int    `json:"up_time"`
	V         int    `json:"__v"`
}

type UpdateWikiReq struct {
	ProjectID   string `json:"project_id"`
	Desc        string `json:"desc"`
	Markdown    string `json:"markdown"`
	EmailNotice bool   `json:"email_notice"`
}
