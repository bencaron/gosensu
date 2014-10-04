package sensu

import "fmt"

// GetEvents Return all the current events
func (s *Sensu) GetEvents() ([]interface{}, error) {
	return s.GetList("events", 0, 0)
}

// GetEventsForClient Returns the current events for given client
func (s *Sensu) GetEventsForClient(client string) ([]interface{}, error) {
	//return s.Get("events", client)
	// TODO  is this the correct way? need validation??
	return s.GetList(fmt.Sprintf("events/%s", client), 0, 0)
}

// GetEventsCheckForClient Returns the event for a check for a client
func (s *Sensu) GetEventsCheckForClient(client string, check string) ([]interface{}, error) {
	//return s.Get("events", client)
	// TODO  is this the correct way? need validation??
	return s.GetList(fmt.Sprintf("events/%s/%s", client, check), 0, 0)
}

// ResolveEvent Resolves an event (delayed action)
func (s *Sensu) ResolveEvent(client string, check string) (map[string]interface{}, error) {
	//return s.Get("events", client)
	fmt.Printf("FIXME ResolveEvent Post: payload missing\n")
	return s.Post(fmt.Sprintf("events/resolve/%s/%s", client, check))
}
