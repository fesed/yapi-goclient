package yapi

import (
	"fmt"

	"github.com/fesed/yapi-goclient/model"
)

func (c *Client) SearchUserByNameOrEmail(key string) ([]*model.User, error) {
	c.refreshLoginStatus()

	var resp model.SearchUserResp
	_, err := c.cli.R().SetBody(&resp).SetResult(&resp).Get(fmt.Sprintf("/api/user/search?q=%s", key))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) IfUserExist(key string) (bool, error) {
	users, err := c.SearchUserByNameOrEmail(key)
	if err != nil {
		return false, err
	}
	for _, user := range users {
		if user.Username == key || user.Email == key {
			return true, nil
		}
	}
	return false, nil
}

func (c *Client) GetUserByNameOrEmail(key string) (*model.User, error) {
	users, err := c.SearchUserByNameOrEmail(key)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == key || user.Email == key {
			return user, nil
		}
	}
	return nil, fmt.Errorf("user not found:%s", key)
}

func (c *Client) GetUserIdByKey(key string) (int, error) {
	user, err := c.GetUserByNameOrEmail(key)
	if err != nil {
		return 0, err
	}
	return user.UID, nil
}

func (c *Client) GetUserIdsByKeys(keys []string) ([]int, error) {
	res := make([]int, len(keys))
	for i, key := range keys {
		id, err := c.GetUserIdByKey(key)
		if err != nil {
			return nil, err
		}
		res[i] = id
	}
	return res, nil
}

// 非官方接口
func (c *Client) AddThirdUser(username, email string) (int, error) {
	c.refreshLoginStatus()

	body := model.AddThirdUserReq{
		Username: username,
		Email:    email,
	}
	var resp model.AddResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/user/add_third_user")
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return 0, err
	}
	return resp.Data.ID, nil
}

func (c *Client) AddThirdUserIfNotExist(username, email string) (int, error) {
	users, err := c.SearchUserByNameOrEmail(username)
	if err != nil {
		return 0, err
	}

	for _, user := range users {
		if user.Username == username {
			return user.UID, nil
		}
	}

	return c.AddThirdUser(username, email)
}
