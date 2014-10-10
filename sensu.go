// Package sensu implements simple methods to interact with Sensu API.
package sensu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	//"bytes"
)

// Sensu struct contains the API details
type Sensu struct {
	Name    string
	Path    string
	URL     string
	Timeout int
	User    string
	Pass    string
}

// NoLimit do not specify a limit parameter
const NoLimit int = -1

// NoOffset do not specify an offset parameter
const NoOffset int = -1

// New Initialize a new Sensu API
func New(name string, path string, url string, timeout int, username string, password string) *Sensu {
	return &Sensu{name, path, url, timeout, username, password}
}

// Health The health endpoint checks to see if the api can connect to redis and rabbitmq. It takes parameters for minimum consumers and maximum messages and checks rabbitmq.
func (s *Sensu) Health(consumers int, messages int) (map[string]interface{}, error) {
	fmt.Printf("FIXME health args")
	return s.Get(fmt.Sprintf("health/%d/%d", consumers, messages))
}

// Info Will return the Sensu version along with rabbitmq and redis information.
func (s *Sensu) Info() (map[string]interface{}, error) {
	return s.Get("info")
}

// Get ...
func (s *Sensu) Get(endpoint string) (map[string]interface{}, error) {
	// Call a List with magic value of limit 0 and offset 0
	url := fmt.Sprintf("%s/%s", s.URL, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Parsing error: %q returned: %v", err, err)
	}

	res, err := s.doHTTP(req)
	if err != nil {
		return nil, fmt.Errorf("API call to %q returned: %v", url, err)
	}
	return s.doJSON(res)
}

// GetList Construct an API call and return the list of results
func (s *Sensu) GetList(endpoint string, limit int, offset int) ([]interface{}, error) {

	/*
		args := ""
		//ERROR GET LIST TODO deal with limit %d and offset %d", limit, offset
		if limit != NOLIMIT {
			args = fmt.Sprintf("%slimit=%d", args, limit)
		}
		if offset != NOOFFSET {
			args = fmt.Sprintf("%soffset=%d", args, limit)
		}
	*/
	url := fmt.Sprintf("%s/%s", s.URL, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("URL Parsing error: %q returned: %v", url, err)
	}
	res, err := s.doHTTP(req)
	if err != nil {
		return nil, fmt.Errorf("API call to %q returned: %v", url, err)
	}
	return s.doJSONArray(res)
}

func (s *Sensu) doHTTP(req *http.Request) ([]byte, error) {

	if s.User != "" && s.Pass != "" {
		req.SetBasicAuth(s.User, s.Pass)
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("%v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Parsing response body returned: %v", err)
	}
	return body, nil
}

// doJsonArray Unmarshall JSON expecting an array
func (s *Sensu) doJSONArray(body []byte) ([]interface{}, error) {
	var results []interface{}
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, fmt.Errorf("Parsing JSON-encoded response body: %v", err)
	}
	return results, nil
}

// doJsonArray Unmarshall JSON expecting a map
func (s *Sensu) doJSON(body []byte) (map[string]interface{}, error) {
	var results map[string]interface{}
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, fmt.Errorf("Parsing JSON-encoded response body: %v", err)
	}
	return results, nil
}

// Post to endpoint
func (s *Sensu) Post(endpoint string) (map[string]interface{}, error) {
	// Call a List with magic value of limit 0 and offset 0

	//ERROR GET LIST TODO deal with limit %d and offset %d", limit, offset

	url := fmt.Sprintf("%s/%s", s.URL, endpoint)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Parsing error: %q returned: %v", url, err)
	}

	res, err := s.doHTTP(req)
	if err != nil {
		return nil, fmt.Errorf("API call to %q returned: %v", url, err)
	}
	return s.doJSON(res)
}

// PostPayload to endpoint
func (s *Sensu) PostPayload(endpoint string, payload string) (map[string]interface{}, error) {
	fmt.Printf("DEBUG: PostPayload = %s\n", payload)
	url := fmt.Sprintf("%s/%s", s.URL, endpoint)

	req, err := http.NewRequest("POST", url, strings.NewReader(fmt.Sprintf("%s\n\n", payload)))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", len(payload)))

	if err != nil {
		return nil, fmt.Errorf("Parsing error: %q returned: %v", url, err)
	}
	res, err := s.doHTTP(req)
	if err != nil {
		return nil, fmt.Errorf("API call to %q returned: %v", url, err)
	}
	return s.doJSON(res)
}

// Delete resource
func (s *Sensu) Delete(endpoint string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", s.URL, endpoint)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Parsing error: %q returned: %v", err, err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API call to %q returned: %v", url, err)
	}

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("%v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("Parsing response body returned: %v", err)
	}
	return s.doJSON(body)
}
