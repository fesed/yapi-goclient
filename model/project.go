package model

type ProjectType string

const (
	ProjectTypePrivate = "private"
	ProjectTypePublic  = "public"
)

type GetProjectResp struct {
	CommonResp
	Data *ProjectDetail `json:"data"`
}

type Env struct {
	Header []interface{} `json:"header"`
	Global []interface{} `json:"global"`
	ID     string        `json:"_id"`
	Name   string        `json:"name"`
	Domain string        `json:"domain"`
}

type ProjectDetail struct {
	SwitchNotice bool          `json:"switch_notice"`
	IsMockOpen   bool          `json:"is_mock_open"`
	Strice       bool          `json:"strice"`
	IsJSON5      bool          `json:"is_json5"`
	ID           int           `json:"_id"`
	Name         string        `json:"name"`
	Basepath     string        `json:"basepath"`
	ProjectType  string        `json:"project_type"`
	UID          int           `json:"uid"`
	GroupID      int           `json:"group_id"`
	Icon         string        `json:"icon"`
	Color        string        `json:"color"`
	AddTime      int           `json:"add_time"`
	UpTime       int           `json:"up_time"`
	Env          []Env         `json:"env"`
	Tag          []interface{} `json:"tag"`
	Cat          []Cat         `json:"cat"`
	Role         string        `json:"role"`
}

type Cat struct {
	Index     int    `json:"index"`
	ID        int    `json:"_id"`
	Name      string `json:"name"`
	ProjectID int    `json:"project_id"`
	Desc      string `json:"desc"`
	UID       int    `json:"uid"`
	AddTime   int    `json:"add_time"`
	UpTime    int    `json:"up_time"`
	V         int    `json:"__v"`
}

type ListProjectResp struct {
	CommonResp
	Data struct {
		List []*Project `json:"list"`
	} `json:"data"`
}

type Project struct {
	SwitchNotice bool   `json:"switch_notice"`
	ID           int    `json:"_id"`
	Name         string `json:"name"`
	Basepath     string `json:"basepath"`
	ProjectType  string `json:"project_type"`
	UID          int    `json:"uid"`
	GroupID      int    `json:"group_id"`
	Icon         string `json:"icon"`
	Color        string `json:"color"`
	AddTime      int    `json:"add_time"`
	UpTime       int    `json:"up_time"`
	Env          []Env  `json:"env"`
	Follow       bool   `json:"follow"`
}

type AddProjectReq struct {
	Name        string      `json:"name"`
	Basepath    string      `json:"basepath"`
	Desc        string      `json:"desc"`
	GroupID     string      `json:"group_id"`
	Icon        string      `json:"icon"`
	Color       string      `json:"color"`
	ProjectType ProjectType `json:"project_type"`
}

type UpdateProjectReq struct {
	Name         string      `json:"name"`
	ProjectType  ProjectType `json:"project_type"`
	Basepath     string      `json:"basepath"`
	SwitchNotice bool        `json:"switch_notice"`
	Desc         string      `json:"desc"`
	ID           int         `json:"id"`
	Env          []Env       `json:"env"`
	GroupID      string      `json:"group_id"`
	Strice       bool        `json:"strice"`
	IsJSON5      bool        `json:"is_json5"`
	Tag          []Tag       `json:"tag"`
}

type Tag struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}
