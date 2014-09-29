package sensu

import (
	"fmt"
	"testing"
)

func getSensuTester() *Sensu {
	sensu := new(Sensu)
	sensu.Name = "Sensu Test API"
	sensu.URL = "http://lpd-sensu-01:4567"
	return sensu
}

func TestInitOject(t *testing.T) {
	//conf := conf.NewConfig("/Users/bcaron/src/go/src/github.com/bencaron/uchiwa/config.json")
	//sensu := Sensu{Name: "testing", URL: "http://localhost:4567"}
	sensu := new(Sensu)
	sensu.Name = "Sensu Test API"
	sensu.URL = "http://localhost:4567"
	if sensu == nil {
		t.Error("Sensu object is nil")
	}
}

func TestSensuTester(t *testing.T) {
	sensu := getSensuTester()
	//sensu.NewSensu()
	if sensu == nil {
		t.Error("Sensu object is nil")
	}
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
