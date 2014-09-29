package sensu

import "fmt"

// GetAggregates Return the list of Aggregates
func (s *Sensu) GetAggregates(limit int, offset int) ([]interface{}, error) {
	return s.GetList("aggregates", limit, offset)
}

// GetAggregate Return Aggregate info
func (s *Sensu) GetAggregate(check string) ([]interface{}, error) {
	return s.Get(fmt.Sprintf("aggregate/%s", check))
}

// GetAggregateIssued Return Aggregate history
func (s *Sensu) GetAggregateIssued(check string, issued string, summarize bool, result bool) ([]interface{}, error) {
	// FIXME
	fmt.Printf("FIXME Aggregate Not handling summarize/result")
	return s.Get(fmt.Sprintf("aggregate/%s/%s", check, issued))
}

// DeleteAggregate Return the list of Aggregates
func (s *Sensu) DeleteAggregate(aggregate string) ([]interface{}, error) {
	return s.Delete(fmt.Sprintf("aggregate/%s", aggregate))
}
