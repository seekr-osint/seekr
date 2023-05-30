package tc

import (
	"fmt"
	"testing"
)

func (test Test[T1, T2]) TcTestHandler(t *testing.T) {
	for _, tc := range test.Cases {
		tc := tc // local copy

		t.Run(fmt.Sprintf("Input: %v", tc.Input), func(t *testing.T) {
			t.Parallel()

			result := test.Func(tc.Input)
			if tc.Expect != result {
				t.Errorf("Expected: %v\nGot: %v\n", tc.Expect, result)
			}
		})
	}
}
