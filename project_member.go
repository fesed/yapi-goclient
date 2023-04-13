package yapi

import (
	"fmt"
	"strconv"

	"github.com/fesed/yapi-goclient/model"

	"github.com/samber/lo"
)

func (c *Client) GetProjectMemberList(projectId int) ([]*model.Member, error) {
	c.refreshLoginStatus()

	var resp model.GetProjectMemberListResp
	_, err := c.cli.R().SetResult(&resp).Get(fmt.Sprintf("/api/project/get_member_list?id=%v", projectId))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) AddProjectMemberByNames(projectId int, usernames []string, role model.RoleType) error {
	uids, err := c.GetUserIdsByKeys(usernames)
	if err != nil {
		return err
	}
	return c.AddProjectMember(projectId, uids, role)
}

func (c *Client) AddProjectMember(projectId int, uids []int, role model.RoleType) error {
	c.refreshLoginStatus()

	stringUids := lo.Map(uids, func(item, _ int) string {
		return strconv.Itoa(item)
	})
	body := model.AddProjectMemberReq{
		ID:         strconv.Itoa(projectId),
		MemberUids: stringUids,
		Role:       role,
	}
	var resp model.CommonResp
	t, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/project/add_member")
	_ = t
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}

func (c *Client) ChangeProjectMemberRole(projectId int, uid int, role model.RoleType) error {
	c.refreshLoginStatus()

	body := model.ChangeProjectMemberRoleReq{
		ID:        projectId,
		MemberUid: strconv.Itoa(uid),
		Role:      role,
	}
	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/project/change_member_role")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteProjectMember(projectId, uid int) error {
	c.refreshLoginStatus()

	body := model.DeleteProjectMemberReq{
		ID:        projectId,
		MemberUID: uid,
	}
	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/project/del_member")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}
