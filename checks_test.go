package sensu

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetChecks(t *testing.T) {
	sensu := getSensuTester()
	res, err := sensu.GetChecks()
	if err != nil {
		t.Error(fmt.Sprintf("Sensu checks returned an error: %v", err))
	}
	if res == nil {
		t.Error("Sensu res is nil")
	}
	//	#fmt.Printf("checks: %v\n", res)

	fmt.Printf("checks[0]: %v\n", res[0])
	command := res[0].(map[string]interface{})
	fmt.Printf("\nchecks[0][command]: %v\n", command["name"])
	fmt.Printf("\nchecks[0][command]: %v\n", res[0].(map[string]interface{})["name"])
	fmt.Printf("\nchecks[0][command].type: %v\n", reflect.TypeOf(res))

}

func TestGetCheck(t *testing.T) {
	sensu := getSensuTester()
	res, err := sensu.GetCheck("check_success")
	if err != nil {
		t.Error(fmt.Sprintf("Sensu GetCheck returned an error: %v", err))
	}
	if res == nil {
		t.Error("Sensu res is nil")
	}
	for k, v := range res {
		fmt.Printf("check: k = %s, v =%v\n", k, v)
	}
	//	#fmt.Printf("checks: %v\n", res)
	//
	// fmt.Printf("checks[0]: %v\n", res[0])
	// command := res[0].(map[string]interface{})
	// fmt.Printf("\nchecks[0][command]: %v\n", command["name"])
	// fmt.Printf("\nchecks[0][command]: %v\n", res[0].(map[string]interface{})["name"])
	// fmt.Printf("\nchecks[0][command].type: %v\n", reflect.TypeOf(res))

}

func TestRequestCheck(t *testing.T) {
	sensu := getSensuTester()

	res, err := sensu.RequestCheck("chef_success")
	if err != nil {
		t.Error(fmt.Sprintf("Sensu checks returned an error: %v", err))
	}
	if res == nil {
		t.Error("Sensu res is nil")
	}
}
