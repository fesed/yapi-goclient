package model

type RoleType string

const (
	// 组长
	RoleOwner RoleType = "owner"
	// 开发者
	RoleDev RoleType = "dev"
	// 访客
	RoleGuest RoleType = "guest"
)

type Member struct {
	ID          string `json:"_id"`
	Role        string `json:"role"`
	UID         int    `json:"uid"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	EmailNotice bool   `json:"email_notice"`
}

type AddGroupMemberReq struct {
	ID         int      `json:"id"`
	MemberUids []string `json:"member_uids"`
	Role       RoleType `json:"role"`
}

type ChangeGroupMemberRoleReq struct {
	ID        int      `json:"id"`
	MemberUid string   `json:"member_uid"`
	Role      RoleType `json:"role"`
}

type GetGroupMemberListResp struct {
	CommonResp
	Data []*Member `json:"data"`
}

type DeleteGroupMemberReq struct {
	ID        int `json:"id"`
	MemberUID int `json:"member_uid"`
}

type GetProjectMemberListResp struct {
	CommonResp
	Data []*Member `json:"data"`
}

type AddProjectMemberReq struct {
	ID         string   `json:"id"`
	MemberUids []string `json:"member_uids"`
	Role       RoleType `json:"role"`
}

type ChangeProjectMemberRoleReq struct {
	ID        int      `json:"id"`
	MemberUid string   `json:"member_uid"`
	Role      RoleType `json:"role"`
}

type DeleteProjectMemberReq struct {
	ID        int `json:"id"`
	MemberUID int `json:"member_uid"`
}
