package model

type LogType string

const (
	LogTypeGroup   LogType = "group"
	LogTypeProject LogType = "project"
)

type GetLogsResp struct {
	CommonResp
	Data struct {
		List  []*Log `json:"list"`
		Total int    `json:"total"`
	} `json:"data"`
}

type Log struct {
	ID       int    `json:"_id"`
	Content  string `json:"content"`
	Type     string `json:"type"`
	UID      int    `json:"uid"`
	Username string `json:"username"`
	Typeid   int    `json:"typeid"`
	AddTime  int    `json:"add_time"`
	V        int    `json:"__v"`
}

type AddUpdateListLogReq struct {
	Type   string `json:"type"`
	Typeid string `json:"typeid"`
	Apis   []struct {
		Method string `json:"method"`
		Path   string `json:"path"`
	} `json:"apis"`
}
