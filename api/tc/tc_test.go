package tc

import (
	"testing"
)
func b(b bool) bool { return b }

func TestNewTest(t *testing.T) {
	tests := map[bool]bool{
		true: true,
		false: false,
	}
	test := NewTest(tests,b)
	test.TcTestHandler(t)
}

