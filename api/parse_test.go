package api

import (
	"reflect"
	"sync"
	"testing"
)

type TestCasePerson struct {
	Input  Person
	expect Person
}
func TcTestHandlerPerson(t *testing.T, testCases []TestCasePerson, testFunc func(Person) Person) { // example TcTestHandler(t,testCases,TestFunction)
	wg := &sync.WaitGroup{}

	for _, tc := range testCases {
		wg.Add(1)
		go func(tc TestCasePerson) {
			result := testFunc(tc.Input)
			if reflect.DeepEqual(tc.expect,result) {
				t.Errorf("Expected %#v for %#v, got %#v", tc.expect, tc.Input, result)
			}
			wg.Done()
		}(tc)
	}
	wg.Wait()
}
func TestParsePerson(t *testing.T) {
	testCases := []TestCasePerson{
    {Person{
      ID: "1",
      Pictures: nil,
      Accounts: nil,
      Sources: nil,
    }, Person{
      ID: "1",
    }},
	}

	TcTestHandlerPerson(t, testCases, ParsePerson)
}

