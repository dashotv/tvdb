package tvdb

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tvdbURL = os.Getenv("TVDB_API_URL")
var tvdbKey = os.Getenv("TVDB_API_KEY")

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

func TestEnv(t *testing.T) {
	assert.NotNil(t, tvdbURL, "TVDB_API_URL is not set")
	assert.NotEmpty(t, tvdbURL, "TVDB_API_URL is empty")

	assert.NotNil(t, tvdbKey, "TVDB_API_KEY is not set")
	assert.NotEmpty(t, tvdbKey, "TVDB_API_KEY is empty")
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

func TestTvdb_Login(t *testing.T) {
	c := New(tvdbURL)
	assert.NotNil(t, c)

	err := c.Login(tvdbKey)
	assert.NoError(t, err)
}

func TestTvdb_Search(t *testing.T) {
	c := New(tvdbURL)
	assert.NotNil(t, c)

	err := c.Login(tvdbKey)
	assert.NoError(t, err)

	p := &GetSearchResultsParams{Query: String("The Simpsons")}
	r, err := c.GetSearchResults(p)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, 38, len(*r.JSON200.Data))
}
