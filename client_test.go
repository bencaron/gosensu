package sensu

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientList(t *testing.T) {
	assert := assert.New(t)
	sensu := getSensuTester()
	events, err := sensu.GetEvents()
	assert.NotNil(events, fmt.Sprintf("Sensu events returned and error: %v", err))
}
