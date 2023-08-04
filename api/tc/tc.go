package tc

import (
	"fmt"
	"strings"
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
func (test Test[T1, T2]) TcTestToMarkdown() string {
	var sb strings.Builder

	for _, testCase := range test.Cases {

		if testCase.Title != "" {
			sb.WriteString(fmt.Sprintf("## **%s**\n\n", testCase.Title))
		} else {
			sb.WriteString("## Test Case\n\n")
		}

		if testCase.Description != "" {
			sb.WriteString(fmt.Sprintf("**Description**: %s\n\n", testCase.Description))
		}

		sb.WriteString("### Input\n\n")
		sb.WriteString(fmt.Sprintf("`%v`\n\n", testCase.Input))

		sb.WriteString("### Output\n\n")
		sb.WriteString(fmt.Sprintf("`%v`\n\n", testCase.Expect))

		//sb.WriteString("### Actual Output\n\n")
		//actual := test.Func(testCase.Input)
		//sb.WriteString(fmt.Sprintf("`%v`\n\n", actual))
	}

	return sb.String()
}
