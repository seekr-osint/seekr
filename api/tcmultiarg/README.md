# tcmultiarg
Test framework used by seekr

## Examples
```go
package main

import (
	"testing"
	"github.com/seekr-osint/seekr/api/tcmultiarg
)

func func1(flag bool, num float64) (string, int) {
	return "hello", int(num)
}

func TestNewTest(t *testing.T) {
	testCases := [][2]tcmultiarg.Args{
		[2]tcmultiarg.Args{tcmultiarg.Args{true, 3.14}, tcmultiarg.Args{"hello", 3}},
		[2]tcmultiarg.Args{tcmultiarg.Args{false, 3.14}, tcmultiarg.Args{"hello", 3}},
	}
	tests := tcmultiarg.NewTest(func1, testCases)
	tests.Run(t)
}
```
