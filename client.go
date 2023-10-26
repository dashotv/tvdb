package tvdb

import (
	"context"

	"github.com/dashotv/tvdb/openapi"
)

type Client struct {
	ApiKey string
	sdk    *openapi.SDK
	ctx    context.Context
}

func New(apikey string) *Client {
	s := openapi.New(openapi.WithSecurity(apikey))
	return &Client{
		ApiKey: apikey,
		sdk:    s,
		ctx:    context.Background(),
	}
}
