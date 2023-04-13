package yapi

import (
	"fmt"

	"github.com/fesed/yapi-goclient/model"
)

func (c *Client) GetSwaggerByUrl(url string) (any, error) {
	c.refreshLoginStatus()

	var resp model.GetSwaggerByUrlResp
	_, err := c.cli.R().SetResult(&resp).Get(fmt.Sprintf("/api/project/swagger_url?url=%s", url))
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (c *Client) UpdateSwaggerSyncConfig(in model.UpdateSwaggerSyncConfigReq) error {
	c.refreshLoginStatus()

	var resp model.CommonResp
	_, err := c.cli.R().SetBody(&in).SetResult(&resp).Post("/api/plugin/autoSync/save")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}
