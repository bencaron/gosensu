package sensu

import (
	"fmt"
	"os"
	"testing"
)

var showedConfig bool = false

func getSensuTester() *Sensu {
	if !showedConfig {
		fmt.Printf("Reading sensu server config from ENV SENSU_SERVER_URL: %s\n", os.Getenv("SENSU_SERVER_URL"))
		showedConfig = true
	}
	sensu := NewSensu("Sensu test API", "", os.Getenv("SENSU_SERVER_URL"), 15, "admin", "secret")
	return sensu
}

func TestSensuTester(t *testing.T) {
	sensu := getSensuTester()
	//sensu.NewSensu()
	if sensu == nil {
		t.Error("Sensu object is nil")
	}
}

func TestSensuInfo(t *testing.T) {
	sensu := getSensuTester()
	if sensu == nil {
		t.Error("Sensu object is nil")
	} /* meh. Need to figure out correct type assertions :/
	info, err := sensu.Info()
	switch ty := info.(type) {
	default:
		t.Error("Sensu object is nil")
	case map[string]interface{}:
		fmt.Printf("sensu.GetInfo() is the correct type")
	}
	*/
}

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
