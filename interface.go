package yapi

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/fesed/yapi-goclient/model"
)

func (c *Client) ListInterface(projectId, page, limit int, status, tag string) ([]*model.Interface, int, error) {
	c.refreshLoginStatus()

	var resp model.ListInterfaceResp
	_, err := c.cli.R().SetResult(&resp).Get(fmt.Sprintf("/api/interface/list?project_id=%v&page=%d&limit=%d&status=%s&tag=%s", projectId, page, limit, status, url.QueryEscape(tag)))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, 0, err
	}
	return resp.Data.List, resp.Data.Count, nil
}

func (c *Client) ListAllInterface(projectId int) ([]*model.Interface, int, error) {
	c.refreshLoginStatus()

	var resp model.ListInterfaceResp
	_, err := c.cli.R().SetResult(&resp).Get(fmt.Sprintf("/api/interface/list?project_id=%v&limit=all", projectId))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, 0, err
	}
	return resp.Data.List, resp.Data.Count, nil
}

func (c *Client) AddInterface(method model.InterfaceMethodType, catId, projectId int, title, path string) (int, error) {
	c.refreshLoginStatus()

	body := model.AddInterfaceReq{
		Method:    method,
		Catid:     strconv.Itoa(catId),
		ProjectID: projectId,
		Path:      path,
		Title:     title,
	}
	var resp model.AddResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/interface/add")
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return 0, err
	}
	return resp.Data.ID, nil
}

func (c *Client) SaveInterface(in *model.SaveInterfaceReq) (int, error) {
	c.refreshLoginStatus()

	var resp model.SaveInterfaceResp
	_, err := c.cli.R().SetBody(in).SetResult(&resp).Post("/api/interface/save")
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return 0, err
	}
	if len(resp.Data) != 0 {
		return resp.Data[0].ID, nil
	}
	return 0, nil
}

func (c *Client) GetInterface(id int) (*model.InterfaceDetail, error) {
	c.refreshLoginStatus()

	var resp model.GetInterfaceResp
	_, err := c.cli.R().SetResult(&resp).Get(fmt.Sprintf("/api/interface/get?id=%v", id))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) UpdateInterface(in *model.UpdateInterfaceReq) error {
	c.refreshLoginStatus()

	var resp model.CommonResp
	_, err := c.cli.R().SetBody(in).SetResult(&resp).Post("/api/interface/up")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteInterface(id int) error {
	c.refreshLoginStatus()

	body := model.DeleteReq{
		Id: id,
	}
	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/interface/del")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}

// 返回分类
func (c *Client) ListMenu(projectId int) ([]*model.Menu, error) {
	c.refreshLoginStatus()

	var resp model.ListMenuResp
	_, err := c.cli.R().SetResult(&resp).Get(fmt.Sprintf("/api/interface/list_menu?project_id=%d", projectId))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// 返回分类下的接口
func (c *Client) ListInterfacesByCat(catId, page, limit int) ([]*model.Interface, int, error) {
	c.refreshLoginStatus()

	var resp model.ListInterfacesByCat
	_, err := c.cli.R().SetResult(&resp).Get(fmt.Sprintf("/api/interface/list_cat?catid=%d&page=%d&limit=%d", catId, page, limit))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, 0, err
	}
	return resp.Data.List, resp.Data.Count, nil
}

func (c *Client) AddCat(name, desc string, projectId int) (int, error) {
	c.refreshLoginStatus()

	body := model.AddCatReq{
		Name:      name,
		Desc:      desc,
		ProjectID: strconv.Itoa(projectId),
	}
	var resp model.AddResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/interface/add_cat")
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return 0, err
	}
	return resp.Data.ID, nil
}

func (c *Client) UpdateCat(id int, name, desc string) error {
	c.refreshLoginStatus()

	body := model.UpdateCatReq{
		Catid: id,
		Name:  name,
		Desc:  desc,
	}
	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("api/interface/up_cat")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteCat(id int) error {
	c.refreshLoginStatus()

	body := model.DeleteCatReq{
		CatId: id,
	}
	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/interface/del_cat")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}
