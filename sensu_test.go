package sensu

import (
  "testing"
  "github.com/bencaron/uchiwa/conf"
)

func TestInitOject(t *testing.T){
  conf := conf.NewConfig("/Users/bcaron/src/go/src/github.com/bencaron/uchiwa/config.json")
  sensu := new(Sensu)
  sensu.New(*conf)
  if sensu == nil {
    t.Error("Sensu object is nil");
  }
}
