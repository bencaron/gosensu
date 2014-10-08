package sensu

import (
	"fmt"
	"testing"
)

func TestGetClientList(t *testing.T) {
	sensu := getSensuTester()
	events, err := sensu.GetEvents()
	if err != nil {
		t.Error(fmt.Sprintf("Sensu events returned and error: %v", err))
	}
	if events == nil {
		t.Error("Sensu events is nil")
	}
	// FIXME
	// fmt.Printf("events: %v", events)
}
