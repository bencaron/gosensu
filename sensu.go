// Package sensu implements simple methods to interact with Sensu API.
package sensu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Sensu struct contains the API details
type Sensu struct {
	Name    string
	Path    string
	URL     string
	Timeout int
}

// Sensu contains data fetched from Sensu API
//var sensu = new(Sensu)

// NewSensu initialize a new Sensu API
func (s *Sensu) NewSensu() {
	fmt.Printf("NewSensu %+v\n", s)
	//s.Config = config
	//fmt.Printf("%+v\n", config)
	//fmt.Printf("%+v\n", s)
	//if s.Name = config.Name; s.Name != "" {
	//  s.Name
	//}
	//s = config
	///fmt.Printf("%+v\n", s)
}

// GetEventsForClient Returns the current events for given client
func (s *Sensu) GetEventsForClient(client string) ([]interface{}, error) {
	//return s.Get("events", client)
	return s.Get(fmt.Sprintf("events/%s", client))
}

// GetEvents Return all the current events
func (s *Sensu) GetEvents() ([]interface{}, error) {
	return s.Get("events")
}

// Get ...
func (s *Sensu) Get(endpoint string) ([]interface{}, error) {
	url := fmt.Sprintf("%s/%s", s.URL, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Sensu.Get URL Parsing error: %q returned: %v", err, err)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request returned: %v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Parsing response body returned: %v", err)
	}

	var results []interface{}
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, fmt.Errorf("Parsing JSON-encoded response body: %v", err)
	}

	return results, nil
}

// func init() {
// 	fmt.Printf("why is this called?")
// 	sensu.Name = "Sensu"
// 	sensu.Path = ""
// 	sensu.Timeout = 10000
// }

//func GetClients(url string) ([]interface{}, error) {
//clients, err := get(url, "clients")
//if err != nil {
//  return nil, fmt.Errorf("%v", err)
//}
//return clients, nil
//else {
//for _, data := range clients {
//  client := data.(map[string]interface{})
//name := client["name"]
//  fmt.Printf("%v\n", client)
//  }
//}
//}
