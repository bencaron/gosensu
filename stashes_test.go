package sensu

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStashes(t *testing.T) {
	assert := assert.New(t)
	sensu := getSensuTester()
	if assert.NotNil(t, sensu) {
		content := make(map[string]interface{})
		content["test"] = "allo"
		stash := Stash{"test/pouelle", content, 30}
		stashes, err := sensu.CreateStash(stash)
		assert.NotNil(stashes, fmt.Sprintf("CreateStash is nil, error is : %s", err))
	}
}

func TestGetStashes(t *testing.T) {
	assert := assert.New(t)
	sensu := getSensuTester()

	if assert.NotNil(t, sensu) {
		stashes, err := sensu.GetStashes()
		assert.NotNil(stashes, fmt.Sprintf("GetStashes is nil, error is : %s", err))
		assert.Equal(stashes, "nope")
	}
}
