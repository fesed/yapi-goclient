package yapi

import (
	"fmt"

	"github.com/fesed/yapi-goclient/model"
)

func (c *Client) GetWiki(projectId int) (*model.Wiki, error) {
	c.refreshLoginStatus()

	var resp model.GetWikiResp
	_, err := c.cli.R().SetResult(&resp).Get(fmt.Sprintf("/api/plugin/wiki_desc/get?project_id=%d", projectId))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) UpdateWiki(in model.UpdateWikiReq) error {
	c.refreshLoginStatus()

	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&in).SetResult(&resp).Post("/api/plugin/wiki_desc/up")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}
