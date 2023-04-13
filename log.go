package yapi

import (
	"fmt"

	"github.com/fesed/yapi-goclient/model"
)

func (c *Client) GetLogs(typ model.LogType, typId, page, limit int, selectValue ...int) ([]*model.Log, int, error) {
	c.refreshLoginStatus()

	path := fmt.Sprintf("/api/log/list?type=%s&typeid=%v&page=%v&limit=%v", typ, typId, page, limit)
	if len(selectValue) != 0 {
		path = fmt.Sprintf("/api/log/list?type=%s&typeid=%v&page=%v&limit=%v&selectValue=%v", typ, typId, page, limit, selectValue[0])
	}
	var resp model.GetLogsResp
	_, err := c.cli.R().SetResult(&resp).Get(path)
	if err := c.HandlerError(resp.CommonResp, err); err != nil {
		return nil, 0, err
	}
	return resp.Data.List, resp.Data.Total, nil
}

func (c *Client) AddUpdateListLog(in *model.AddUpdateListLogReq) error {
	c.refreshLoginStatus()

	var resp model.CommonResp
	_, err := c.cli.R().SetBody(in).SetResult(&resp).Post("/api/log/list_by_update")
	if err := c.HandlerError(resp, err); err != nil {
		return err
	}
	return nil
}
