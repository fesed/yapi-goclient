package model

type CreateGroupReq struct {
	GroupName string   `json:"group_name"`
	GroupDesc string   `json:"group_desc"`
	OwnerUids []string `json:"owner_uids"`
}

type UpdateGroupReq struct {
	GroupName    string `json:"group_name"`
	GroupDesc    string `json:"group_desc"`
	CustomField1 struct {
		Name   string `json:"name"`
		Enable bool   `json:"enable"`
	} `json:"custom_field1"`
	ID int `json:"id"`
}

type CreateGroupResp struct {
	CommonResp
	Data struct {
		ID int `json:"_id"`
	} `json:"data"`
}

type ListGroupResp struct {
	CommonResp
	Data []*Group `json:"data"`
}

type CustomField1 struct {
	Enable bool `json:"enable"`
}

type Group struct {
	UID       int        `json:"uid,omitempty"`
	GroupName string     `json:"group_name"`
	Type      string     `json:"type"`
	ID        int        `json:"_id"`
	GroupDesc string     `json:"group_desc,omitempty"`
	Role      string     `json:"role,omitempty"`
	Members   []*Members `json:"members"`
}

type Members struct {
	ID       string `json:"_id"`
	Role     string `json:"role"`
	UID      int    `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
