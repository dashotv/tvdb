package tvdb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tvdbApiKey = os.Getenv("TVDB_API_KEY")

func TestClient_Login(t *testing.T) {
	assert.NotEmpty(t, tvdbApiKey, "TVDB_API_KEY is empty")

	c := New(tvdbApiKey)

	token, err := c.Login()
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// fmt.Printf("Token: %s\n", token)
}
