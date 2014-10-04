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
	Pass    string
	Path    string
	URL     string
	User    string
	Timeout int
}

// NewSensu initialize a new Sensu API
func NewSensu(name string, pass string, path string, url string, user string, timeout int) *Sensu {
	return &Sensu{name, pass, path, url, user, timeout}
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

	//ERROR GET LIST TODO deal with limit %d and offset %d", limit, offset

	url := fmt.Sprintf("%s/%s", s.URL, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Sensu.Get URL Parsing error: %q returned: %v", err, err)
	}

	if s.User != "" && s.Pass != "" {
		req.SetBasicAuth(s.User, s.Pass)
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

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("%s: %s", url, res.Status)
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
