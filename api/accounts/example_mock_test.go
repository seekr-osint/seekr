package accounts_test

import (
	"fmt"

	"github.com/seekr-osint/seekr/api/accounts"
)

func ExampleGetFilePath() {
	path := accounts.GetFilePath("https://github.com/greg")
	fmt.Printf("%v", path)
	// Output:
	//
	// mock/aHR0cHM6Ly9naXRodWIuY29tL2dyZWc=.json
}
