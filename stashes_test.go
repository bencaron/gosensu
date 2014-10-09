package sensu

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testPath string = "pouelle"

func TestCreateStashes(t *testing.T) {
	assert := assert.New(t)
	sensu := getSensuTester()
	if assert.NotNil(t, sensu) {
		content := make(map[string]interface{})
		content["test"] = "allo"
		stash := Stash{testPath, content, 30}
		stashes, err := sensu.CreateStash(stash)
		fmt.Printf("\nstashes: %v err %s\n", stashes, err)
		assert.NotNil(stashes, fmt.Sprintf("CreateStash is nil, error is : %s", err))
	}
}

func TestCreateStashePath(t *testing.T) {
  assert := assert.New(t)
  sensu := getSensuTester()
  if assert.NotNil(t, sensu) {
    content := make(map[string]interface{})
    content["test"] = "allo"
    stashes, err := sensu.CreateStashPath("testPath2", content)
    fmt.Printf("\nstashes: %v err %s\n", stashes, err)
    assert.NotNil(stashes, fmt.Sprintf("CreateStash is nil, error is : %s", err))
  }
}

func TestGetStashes(t *testing.T) {
	assert := assert.New(t)
	sensu := getSensuTester()
	var empty []interface{}
	if assert.NotNil(t, sensu) {
		stashes, err := sensu.GetStashes()
		assert.NotNil(stashes, fmt.Sprintf("GetStashes is nil, error is : %s", err))
		assert.Equal(stashes, empty)
	}
}
