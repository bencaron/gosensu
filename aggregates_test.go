package sensu

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetAggregates(t *testing.T) {
	sensu := getSensuTester()
	if strings.Contains(sensu.URL, "localhost") {
		t.Skip("Not testing aggregates with canned")
	}
	agg, err := sensu.GetAggregates()
	if err != nil {
		t.Error(fmt.Sprintf("Sensu agg returned and error: %v", err))
	}
	if agg == nil {
		t.Error("Sensu agg is nil")
	}
	fmt.Printf("TestGet Aggregates:\n")
	fmt.Printf("\t\tagg: %v", agg)
}
