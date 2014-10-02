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
	agg, err := sensu.GetAggregates(10, 0)
	if err != nil {
		t.Error(fmt.Sprintf("Sensu agg returned and error: %v", err))
	}
	if agg == nil {
		t.Error("Sensu agg is nil")
	}
	//fmt.Printf("agg: %v", agg)
}
