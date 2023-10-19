package tvdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testClient(t *testing.T) *ClientWithResponses {
	c, err := newClientWithResponses(tvdbURL)
	assert.NoError(t, err)
	assert.NotNil(t, c)

	return c
}

func testAuthClient(t *testing.T) (*ClientWithResponses, context.Context) {
	c := testClient(t)
	ctx := context.Background()

	b := PostLoginJSONRequestBody{}
	b.Apikey = tvdbKey

	resp, err := c.PostLoginWithResponse(ctx, b)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	fmt.Printf("Token: %s\n", *resp.JSON200.Data.Token)

	ctx = context.WithValue(ctx, contextBearerToken, *resp.JSON200.Data.Token)

	nc, err := newClientWithResponses(tvdbURL, WithRequestEditorFn(requestAuth))
	assert.NoError(t, err)
	assert.NotNil(t, nc)

	return nc, ctx
}

func TestClient_Login(t *testing.T) {
	c := testClient(t)
	ctx := context.Background()

	b := PostLoginJSONRequestBody{}
	b.Apikey = tvdbKey

	resp, err := c.PostLoginWithResponse(ctx, b)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	fmt.Printf("Token: %s\n", *resp.JSON200.Data.Token)
}

func TestClient_Search(t *testing.T) {
	c, ctx := testAuthClient(t)

	q := "The Simpsons"
	b := &GetSearchResultsParams{Query: &q}

	resp, err := c.GetSearchResultsWithResponse(ctx, b)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	fmt.Printf("Search Results: %d\n", len(*resp.JSON200.Data))
}
