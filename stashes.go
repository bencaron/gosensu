package sensu

import "fmt"

// GetStashes Return a list of stashes path
func (s *Sensu) GetStashes() ([]interface{}, error) {
	return s.GetList(fmt.Sprintf("stashes"), 0, 0)
}

// GetStashesSlice Return a slice in the list of stashes path
func (s *Sensu) GetStashesSlice(limit int, offset int) ([]interface{}, error) {
	fmt.Printf("FIXME GetStashes incomplete %d %d", limit, offset)
	return s.GetList(fmt.Sprintf("stashes"), 0, 0)
}

// GetStash Get a stash
func (s *Sensu) GetStash(path string) (map[string]interface{}, error) {
	return s.Get(fmt.Sprintf("stashes/%s", path))
}

// CreateStash create a stash (JSON document)
func (s *Sensu) CreateStash(check string, payload interface{}) (map[string]interface{}, error) {
	fmt.Printf("FIXME createStash POST %v\n", payload)
	//	return s.Post(fmt.Sprintf("stashes/create"), payload)
	return s.PostPayload(fmt.Sprintf("stashes/create"), "{  \"broken\": true}")
}

// DeleteStash Delete a stash (JSON document)
func (s *Sensu) DeleteStash(path string) (map[string]interface{}, error) {
	return s.Delete(fmt.Sprintf("stashes/%s", path))
}
