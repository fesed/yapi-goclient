package yapi

import (
	"fmt"
	"strconv"

	"github.com/fesed/yapi-goclient/model"

	"github.com/samber/lo"
)

func (c *Client) GetGroupMemberList(groupId int) ([]*model.Member, error) {
	c.refreshLoginStatus()

	var resp model.GetGroupMemberListResp
	_, err := c.cli.R().SetResult(&resp).Get(fmt.Sprintf("/api/group/get_member_list?id=%v", groupId))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) AddGroupMemberByNames(groupId int, usernames []string, role model.RoleType) error {
	uids, err := c.GetUserIdsByKeys(usernames)
	if err != nil {
		return err
	}

	return c.AddGroupMember(groupId, uids, role)
}

func (c *Client) AddGroupMember(groupId int, uids []int, role model.RoleType) error {
	c.refreshLoginStatus()

	stringUids := lo.Map(uids, func(item, _ int) string {
		return strconv.Itoa(item)
	})
	body := model.AddGroupMemberReq{
		ID:         groupId,
		MemberUids: stringUids,
		Role:       role,
	}
	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/group/add_member")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}

func (c *Client) ChangeGroupMemberRole(groupId int, uid int, role model.RoleType) error {
	c.refreshLoginStatus()

	body := model.ChangeGroupMemberRoleReq{
		ID:        groupId,
		MemberUid: strconv.Itoa(uid),
		Role:      role,
	}
	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/group/change_member_role")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteGroupMember(groupId, uid int) error {
	c.refreshLoginStatus()

	body := model.DeleteGroupMemberReq{
		ID:        groupId,
		MemberUID: uid,
	}
	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/group/del_member")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}
