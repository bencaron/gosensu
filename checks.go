package sensu

import "fmt"

// GetChecks Return the list of checks
func (s *Sensu) GetChecks() ([]interface{}, error) {
	return s.Get("checks")
}

// GetCheck Return check info
func (s *Sensu) GetCheck(check string) ([]interface{}, error) {
	return s.Get(fmt.Sprintf("check/%s", check))
}

// RequestCheck Issues a check request
func (s *Sensu) RequestCheck(check string) ([]interface{}, error) {
	return s.Post(fmt.Sprintf("check/request"))
}
