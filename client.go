package yapi

import (
	"errors"

	"github.com/fesed/yapi-goclient/model"

	"github.com/go-resty/resty/v2"
)

type Options struct {
	Url      string `mapstructure:"url"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Client struct {
	opts Options
	cli  *resty.Client
}

func NewClient(opts Options) (*Client, error) {
	client := resty.New()
	client.SetBaseURL(opts.Url)
	client.SetRetryCount(1)

	return &Client{
		opts: opts,
		cli:  client,
	}, nil
}

func (c *Client) GetOpts() Options {
	return c.opts
}

func (c *Client) HandlerError(resp model.CommonResp, err error) error {
	if err != nil {
		return err
	}
	if resp.Errcode != model.StatusOK {
		return errors.New(resp.Errmsg)
	}
	return nil
}
