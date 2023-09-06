package accounts_test

import (
	"fmt"

	"github.com/seekr-osint/seekr/api/accounts"
)

func ExampleParseCheckResult_true() {
	exists, rateLimited, err := accounts.ParseCheckResult("true")
	fmt.Printf("%v, %v, %v", exists, rateLimited, err)
	// Output:
	//
	// true, false, <nil>
}

func ExampleParseCheckResult_false() {
	exists, rateLimited, err := accounts.ParseCheckResult("false")
	fmt.Printf("%v, %v, %v", exists, rateLimited, err)
	// Output:
	//
	// false, false, <nil>
}

// A string starting with `error: ` followed by a error message will return the error message as err.
func ExampleParseCheckResult_err() {
	exists, rateLimited, err := accounts.ParseCheckResult("error: useful err message")
	fmt.Printf("%v, %v, %v", exists, rateLimited, err)
	// Output:
	//
	// false, false, useful err message
}


func ExampleParseCheckResult_errEmpty() {
	exists, rateLimited, err := accounts.ParseCheckResult("error: ")
	fmt.Printf("%v, %v, %v", exists, rateLimited, err)
	// Output:
	//
	// false, false, error message missing
}

func ExampleParseCheckResult_unknwnResult() {
	exists, rateLimited, err := accounts.ParseCheckResult("kinda true")
	fmt.Printf("%v, %v, %v", exists, rateLimited, err)
	// Output:
	//
	// false, false, unknown check result: kinda true
}

func ExampleParseCheckResult_empty() {
	exists, rateLimited, err := accounts.ParseCheckResult("")
	fmt.Printf("%v, %v, %v", exists, rateLimited, err)
	// Output:
	//
	// false, false, empty result
}
