package sensu

import "fmt"

// GetStashes Return a list of stashes path
func (s *Sensu) GetStashes(limit int, offset int) ([]string, error) {
	fmt.Printf("FIXME GetStashes incomplete %d %d", limit, offset)
	//return s.Get("stashes")
	//return ["FIXME", "TEST"], err
	out := make([]string, 2)
	return out, nil
}

// GetStash Get a stash
func (s *Sensu) GetStash(path string) ([]interface{}, error) {
	return s.Get(fmt.Sprintf("stashes/%s", path))
}

// CreateStash create a stash (JSON document)
func (s *Sensu) CreateStash(check string, payload interface{}) ([]interface{}, error) {
	fmt.Printf("FIXME createStash POST %v\n", payload)
	//	return s.Post(fmt.Sprintf("stashes/create"), payload)
	return s.Post(fmt.Sprintf("stashes/create"))
}

// DeleteStash Delete a stash (JSON document)
func (s *Sensu) DeleteStash(path string) ([]interface{}, error) {
	return s.Delete(fmt.Sprintf("stashes/%s", path))
}
