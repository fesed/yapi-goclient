package yapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/fesed/yapi-goclient/model"
)

const (
	freshCycle = 12 * 60 * 60
)

// 如果cookie为空或者cookie中的token过期时间小于12小时，刷新token
// 否则，直接返回
func (c *Client) refreshLoginStatus() error {
	if len(c.cli.Cookies) != 0 {
		for _, cookie := range c.cli.Cookies {
			if cookie.Name == "_yapi_token" {
				cs, err := getClaims(cookie.Value)
				if err != nil {
					return err
				}
				if time.Now().Unix()+freshCycle < int64(cs.Exp) {
					return nil
				}
			}
		}
	}
	return c.login()
}

func (c *Client) login() error {
	body := map[string]string{
		"email":    c.opts.Username,
		"password": c.opts.Password,
	}
	res, err := c.cli.R().SetBody(body).Post("/api/user/login")
	if err != nil {
		return err
	}
	c.cli.SetCookies(res.Cookies())
	return nil
}

func getClaims(jwtToken string) (model.Claims, error) {
	var cs model.Claims
	parts := strings.Split(jwtToken, ".")
	if len(parts) != 3 {
		return cs, fmt.Errorf("compact JWS format must have three parts")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return cs, err
	}

	err = json.Unmarshal(payload, &cs)
	if err != nil {
		return cs, err
	}

	return cs, nil
}
