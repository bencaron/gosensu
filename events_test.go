package sensu

import (
	"fmt"
	"testing"
)

func TestEvents(t *testing.T) {
	fmt.Printf("Test Events:\n*******************\n\n")
}

func TestGetEvents(t *testing.T) {
	fmt.Printf("Get Events:\n")
	sensu := getSensuTester()
	events, err := sensu.GetEvents()
	if err != nil {
		t.Error(fmt.Sprintf("Sensu events returned and error: %v", err))
	}
	if events == nil {
		t.Error("Sensu events is nil")
	}
	fmt.Printf("TestGet Events:\n")
	fmt.Printf("\t\tev: %v", events)
}

func TestResolveEvents(t *testing.T) {
	fmt.Printf("ResolveEvents:\n")
	sensu := getSensuTester()
	events, err := sensu.ResolveEvent("server-0-13-0", "check_critical")
	if err != nil {
		t.Error(fmt.Sprintf("Sensu events returned and error: %v\n", err))
	}
	fmt.Printf("\treturned:\n\t\t%v\n\n", events)
}

func TestResolveNonExistingEvents(t *testing.T) {
	fmt.Printf("ResolveEvents (non-existing):\n")
	sensu := getSensuTester()
	ev, err := sensu.ResolveEvent("server-0-13-0", "check_not_real")
	if err == nil {
		t.Error(fmt.Sprintf("Sensu Resolve Events should not returned and error on non-existing checks: %v", ev))
	}
}
