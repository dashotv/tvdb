package tvdb

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dashotv/tvdb/openapi/models/operations"
)

var tvdbApiKey = os.Getenv("TVDB_API_KEY")
var tvdbToken = os.Getenv("TVDB_API_TOKEN")

func TestLogin(t *testing.T) {
	assert.NotEmpty(t, tvdbApiKey, "TVDB_API_KEY is empty")

	c, err := Login(tvdbApiKey)
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestNew(t *testing.T) {
	assert.NotEmpty(t, tvdbToken, "TVDB_API_TOKEN is empty")

	c := New(tvdbApiKey, tvdbToken)
	assert.NotNil(t, c)
}

func TestClient_Search(t *testing.T) {
	assert.NotEmpty(t, tvdbApiKey, "TVDB_API_KEY is empty")

	c, err := Login(tvdbApiKey)
	assert.NoError(t, err)
	assert.NotNil(t, c)

	p := operations.GetSearchResultsRequest{
		Query: String("The Simpsons"),
		Type:  String("series"),
	}
	r, err := c.GetSearchResults(p)
	assert.NoError(t, err)
	assert.NotNil(t, r)

	for _, v := range r.Data {
		assert.NotEmpty(t, *v.ID)
		assert.NotEmpty(t, *v.Name)
		fmt.Printf("%s: %s\n", *v.ID, *v.Name)
	}
}
