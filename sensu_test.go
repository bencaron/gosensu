package sensu

import (
	"fmt"
	"os"
	"testing"
)

func getSensuTester() *Sensu {
	fmt.Printf("Reading sensu server config from ENV SENSU_SERVER_URL: %s\n", os.Getenv("SENSU_SERVER_URL"))
	sensu := NewSensu("Sensu test API", "", os.Getenv("SENSU_SERVER_URL"), 15, "admin", "secret")
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
