package tcmultiarg

import (
	"testing"
)

func func1(flag bool, num float64) (string, int) {
	return "hello", int(num)
}

func TestNewTest(t *testing.T) {
	testCases := [][2]Args{
		[2]Args{Args{true, 3.14}, Args{"hello", 3}},
		[2]Args{Args{false, 3.14}, Args{"hello", 3}},
	}
	tests := NewTest(func1, testCases)
	tests.Run(t)
}
