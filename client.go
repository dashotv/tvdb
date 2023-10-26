package tvdb

import (
	"context"

	"github.com/dashotv/tvdb/openapi"
	"github.com/dashotv/tvdb/openapi/models/operations"
)

type Client struct {
	Key   string
	Token string
	sdk   *openapi.SDK
	ctx   context.Context
}

func New(apikey string) *Client {
	s := openapi.New(openapi.WithSecurity(apikey))
	return &Client{
		Key:   apikey,
		Token: "",
		sdk:   s,
		ctx:   context.Background(),
	}
}

func (c *Client) Login() (string, error) {
	p := operations.PostLoginRequestBody{
		Apikey: c.Key,
	}

	r, err := c.PostLogin(p)
	if err != nil {
		return "", err
	}

	c.Token = *r.Data.Token
	return c.Token, nil
}
