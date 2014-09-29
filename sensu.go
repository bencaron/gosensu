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

// Health The health endpoint checks to see if the api can connect to redis and rabbitmq. It takes parameters for minimum consumers and maximum messages and checks rabbitmq.
func (s *Sensu) Health(consumers int, messages int) ([]interface{}, error) {
	fmt.Printf("FIXME health args")
	return s.Get(fmt.Sprintf("health/%d/%d", consumers, messages))
}

// Info Will return the Sensu version along with rabbitmq and redis information.
func (s *Sensu) Info() ([]interface{}, error) {
	return s.Get("info")
}

// Delete Delete resource
func (s *Sensu) Delete(endpoint string) ([]interface{}, error) {
	// Call a List with magic value of limit 0 and offset 0
	fmt.Printf("FIXME Delete is NOT IMPLEMENTED")
	return s.GetList(endpoint, 0, 0)
}

// Get ...
func (s *Sensu) Get(endpoint string) ([]interface{}, error) {
	// Call a List with magic value of limit 0 and offset 0
	return s.GetList(endpoint, 0, 0)
}

// GetList Construct an API call and return the list of results
func (s *Sensu) GetList(endpoint string, limit int, offset int) ([]interface{}, error) {

	fmt.Printf("ERROR GET LIST TODO deal with limit %d and offset %d", limit, offset)

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

// Post to endpoint
func (s *Sensu) Post(endpoint string) ([]interface{}, error) {
	// Call a List with magic value of limit 0 and offset 0
	fmt.Printf("POST is not implemented yet\n")
	return s.Get(endpoint)
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
