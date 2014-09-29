package sensu

import "fmt"

// GetEvents Return all the current events
func (s *Sensu) GetEvents() ([]interface{}, error) {
	return s.Get("events")
}

// GetEventsForClient Returns the current events for given client
func (s *Sensu) GetEventsForClient(client string) ([]interface{}, error) {
	//return s.Get("events", client)
	// TODO  is this the correct way? need validation??
	return s.Get(fmt.Sprintf("events/%s", client))
}

// GetEventsCheckForClient Returns the event for a check for a client
func (s *Sensu) GetEventsCheckForClient(client string, check string) ([]interface{}, error) {
	//return s.Get("events", client)
	// TODO  is this the correct way? need validation??
	return s.Get(fmt.Sprintf("events/%s/%s", client, check))
}

// ResolveEvent Resolves an event (delayed action)
func (s *Sensu) ResolveEvent(client string, check string) ([]interface{}, error) {
	//return s.Get("events", client)
	fmt.Printf("FIXME ResolveEvent Post: payload missing\n")
	return s.Post(fmt.Sprintf("events/resolve/%s/%s", client, check))
}
