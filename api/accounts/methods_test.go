package accounts

// import (
// 	// "fmt"
// 	"log"
// 	"testing"

// 	"github.com/seekr-osint/seekr/api/tcmultiarg"
// )

// func RunUserExistCheckInputAndAccountScanner(accountScanner AccountScanner, username string) (bool,bool, error) {
// 	userExistsCheckInput, a, err := GetUserExistCheckInputAndAccountScanner(accountScanner, username)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return userExistsCheckInput.RunUserExistsCheck(a)
// }
// func GetUserExistCheckInputAndAccountScanner(accountScanner AccountScanner, username string) (*UserExistsCheckInput, *AccountScanner, error) {
// 	userExistsCheckInput, err := accountScanner.UserExistsCheckInput(User{
// 		Username: username,
// 	})
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	return userExistsCheckInput, &accountScanner, nil
// }

// func TestRunCheck(t *testing.T) {
// 	testCases := [][2]tcmultiarg.Args{
// 		[2]tcmultiarg.Args{tcmultiarg.Args{DefaultServices[0], "greg"}, tcmultiarg.Args{true, nil}},
// 		// [2]tcmultiarg.Args{tcmultiarg.Args{DefaultServices[0], "greg"}, tcmultiarg.Args{false, nil}},
// 		// [2]tcmultiarg.Args{tcmultiarg.Args{"false"}, tcmultiarg.Args{false, nil}},
// 		// [2]tcmultiarg.Args{tcmultiarg.Args{"error: very useful msg"}, tcmultiarg.Args{false, fmt.Errorf("very useful msg")}},
// 		// [2]tcmultiarg.Args{tcmultiarg.Args{"rate limited"}, tcmultiarg.Args{false, fmt.Errorf("rate limited")}},
// 		// [2]tcmultiarg.Args{tcmultiarg.Args{"kinda true", ""}, tcmultiarg.Args{false, fmt.Errorf("unknown check result: %s", "kinda true")}},
// 		// [2]tcmultiarg.Args{tcmultiarg.Args{""}, tcmultiarg.Args{false, fmt.Errorf("empty result")}},
// 	}
// 	tests := tcmultiarg.NewTest(RunUserExistCheckInputAndAccountScanner, testCases)
// 	tests.Run(t)
// }
