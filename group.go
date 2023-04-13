package yapi

import (
	"errors"
	"strconv"

	"github.com/fesed/yapi-goclient/model"

	"github.com/samber/lo"
)

func (c *Client) ListGroup() ([]*model.Group, error) {
	c.refreshLoginStatus()

	var resp model.ListGroupResp
	_, err := c.cli.R().SetResult(&resp).Get("/api/group/list")
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) GetGroupById(id int) (*model.Group, error) {
	groups, err := c.ListGroup()
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		if group.ID == id {
			return group, nil
		}
	}
	return nil, nil
}

func (c *Client) GetGroupByName(name string) (*model.Group, error) {
	groups, err := c.ListGroup()
	if err != nil {
		return nil, err
	}
	for _, group := range groups {
		if group.GroupName == name {
			return group, nil
		}
	}
	return nil, nil
}

func (c *Client) AddGroupWithOwnernames(name, desc string, ownernames ...string) (int, error) {
	userIds, err := c.GetUserIdsByKeys(ownernames)
	if err != nil {
		return 0, err
	}
	userStringIds := lo.Map(userIds, func(item, _ int) string {
		return strconv.Itoa(item)
	})
	return c.AddYapiGroup(name, desc, userStringIds...)
}

func (c *Client) AddYapiGroup(name, desc string, userIds ...string) (int, error) {
	c.refreshLoginStatus()

	body := model.CreateGroupReq{
		GroupName: name,
		GroupDesc: desc,
		OwnerUids: userIds,
	}
	var resp model.CreateGroupResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/group/add")
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return 0, err
	}
	return resp.Data.ID, nil
}

func (c *Client) UpdateGroup(in model.UpdateGroupReq) error {
	c.refreshLoginStatus()

	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&in).SetResult(&resp).Post("/api/group/add")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetGroupIdByName(key string) (int, error) {
	groups, err := c.ListGroup()
	if err != nil {
		return 0, err
	}

	for _, group := range groups {
		if group.GroupName == key {
			return group.ID, nil
		}
	}
	return 0, errors.New("没找到group:" + key)
}

func (c *Client) DeleteGroupByName(groupname string) error {
	groupId, err := c.GetGroupIdByName(groupname)
	if err != nil {
		return err
	}
	return c.DeleteGroup(groupId)
}

func (c *Client) DeleteGroup(groupId int) error {
	c.refreshLoginStatus()

	body := model.DeleteReq{
		Id: groupId,
	}
	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/group/del")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}
