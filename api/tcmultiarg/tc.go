package tcmultiarg

import (
	"reflect"
	"testing"
)

func NewTest(fn interface{}, testCases [][2]Args) Tests {
	tests := Tests{}
	for _, testCase := range testCases {
		if len(testCase) == 2 {
			test := Test{
				Input:    testCase[0],
				Expect:   testCase[1],
				Function: fn,
			}
			tests = append(tests, test)

		} else {
			// error
		}
	}
	return tests
}
func (tests Tests) Run(t *testing.T) bool {
	for _, test := range tests {
		test.Run(t)
	}
	return true
}

func (test Test) Run(t *testing.T) {
	res := ExtractReturnValues(test.Function, test.Input)
	if !reflect.DeepEqual(res, test.Expect) {
		t.Errorf("Expected: %v Got %v", test.Expect, res)
	}
}

func ExtractReturnValues(fn interface{}, args Args) Args {
	fnValue := reflect.ValueOf(fn)
	argValues := make([]reflect.Value, len(args))

	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}

	results := fnValue.Call(argValues)

	var returnValues []interface{}
	for _, result := range results {
		returnValues = append(returnValues, result.Interface())
	}
	return returnValues
}
