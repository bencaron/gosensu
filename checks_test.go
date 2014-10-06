package sensu

import (
	"fmt"
	"testing"
)

func TestChecks(t *testing.T) {
	fmt.Printf("Test Checks:\n*******************\n\n")
}

func TestGetChecks(t *testing.T) {
	fmt.Printf("\tGet Checks:\n")
	sensu := getSensuTester()
	res, err := sensu.GetChecks()
	if err != nil {
		t.Error(fmt.Sprintf("Sensu checks returned an error: %v", err))
	}
	if res == nil {
		t.Error("Sensu res is nil")
	}
	fmt.Printf("\tGet Checks got:\n\t\t%v\n", res)

	/*
		fmt.Printf("\tchecks[0]: %v\n", res[0])
		command := res[0].(map[string]interface{})
		fmt.Printf("\t\nchecks[0][command]: %v\n", command["name"])
		fmt.Printf("\t\nchecks[0][command]: %v\n", res[0].(map[string]interface{})["name"])
		fmt.Printf("\t\nchecks[0][command].type: %v\n", reflect.TypeOf(res))
	*/
}

func TestGetCheck(t *testing.T) {
	fmt.Printf("\tGet (ONE) Check:\n")
	sensu := getSensuTester()
	res, err := sensu.GetCheck("check_success")
	if err != nil {
		t.Error(fmt.Sprintf("Sensu GetCheck returned an error: %v", err))
	}
	if res == nil {
		t.Error("Sensu res is nil")
	}
	/*for k, v := range res {
		fmt.Printf("check: k = %s, v =%v\n", k, v)
	}*/
	fmt.Printf("\tTest Get (one) Check result: %v\n\t\t", res)
	//
	// fmt.Printf("checks[0]: %v\n", res[0])
	// command := res[0].(map[string]interface{})
	// fmt.Printf("\nchecks[0][command]: %v\n", command["name"])
	// fmt.Printf("\nchecks[0][command]: %v\n", res[0].(map[string]interface{})["name"])
	// fmt.Printf("\nchecks[0][command].type: %v\n", reflect.TypeOf(res))

}

func TestRequestCheck(t *testing.T) {
	fmt.Printf("\tRequest Check:\n")
	sensu := getSensuTester()

	t.Skip("Skipping TestRequestCheck, code not ready yet.")

	res, err := sensu.RequestCheck("chef_success")
	if err != nil {
		t.Error(fmt.Sprintf("Sensu checks returned an error: %v", err))
	}
	if res == nil {
		t.Error("Sensu res is nil")
	}
}
