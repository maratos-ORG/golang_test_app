package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBackendParameters(t *testing.T) {
	// Save current os.Args
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "--api-port=9000", "--log-level=DEBUG", "--db_conn=test_connection"}

	params, err := GetBackendParameters()
	assert.NoError(t, err)
	assert.Equal(t, "9000", *params.Port)
	assert.Equal(t, "DEBUG", *params.LogLevel)
	assert.Equal(t, "test_connection", *params.DBURL)
}

// explanation https://chat.openai.com/share/4af3b1e8-38f6-4b49-b57c-c8a8a35b9a3f