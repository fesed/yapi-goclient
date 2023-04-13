package yapi

import (
	"fmt"
	"strconv"

	"github.com/fesed/yapi-goclient/model"
)

func (c *Client) ListProject(groupId int) ([]*model.Project, error) {
	c.refreshLoginStatus()

	var resp model.ListProjectResp
	_, err := c.cli.R().SetResult(&resp).Get(fmt.Sprintf("/api/project/list?group_id=%v", groupId))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, err
	}
	return resp.Data.List, nil
}

func (c *Client) GetProjectByName(groupId int, projectname string) (*model.Project, error) {
	projects, err := c.ListProject(groupId)
	if err != nil {
		return nil, err
	}
	for _, item := range projects {
		if item.Name == projectname {
			return item, nil
		}
	}
	return nil, nil
}

// 如果id不存在，yapi会直接报“系统错误”
func (c *Client) GetProject(projectId int) (*model.ProjectDetail, error) {
	c.refreshLoginStatus()

	var resp model.GetProjectResp
	_, err := c.cli.R().SetResult(&resp).Get(fmt.Sprintf("/api/project/get?id=%v", projectId))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) AddProject(name, desc, basePath string, groupId int) (int, error) {
	c.refreshLoginStatus()

	body := model.AddProjectReq{
		Name:        name,
		Basepath:    basePath,
		Desc:        desc,
		GroupID:     strconv.Itoa(groupId),
		Icon:        "code-o",
		Color:       "blue",
		ProjectType: model.ProjectTypePrivate,
	}
	var resp model.AddResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/project/add")
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return 0, err
	}
	return resp.Data.ID, nil
}

func (c *Client) UpdateProject(in model.UpdateProjectReq) error {
	c.refreshLoginStatus()

	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&in).SetResult(&resp).Post("/api/project/up")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteProject(projectId int) error {
	c.refreshLoginStatus()

	body := model.DeleteReq{
		Id: projectId,
	}
	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&body).SetResult(&resp).Post("/api/project/del")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}
