package language

import (
	"reflect"
	"testing"
)

func TestExtractComments(t *testing.T) {
	code := `// This is a single-line comment
/* This is a
   multi-line comment */
# This is a hash comment`

	tests := []struct {
		commentType []CommentType
		expected    []string
	}{
		{
			commentType: []CommentType{DoubleSlash},
			expected:    []string{" This is a single-line comment"},
		},
		{
			commentType: []CommentType{DoubleSlashMultiLine},
			expected:    []string{" This is a\n   multi-line comment "},
		},
		{
			commentType: []CommentType{Hash},
			expected:    []string{" This is a hash comment"},
		},
		{
			commentType: []CommentType{DoubleSlash, DoubleSlashMultiLine},
			expected: []string{
				" This is a single-line comment",
				" This is a\n   multi-line comment ",
			},
		},
	}

	for _, test := range tests {
		result := ExtractComments(code, test.commentType...)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ExtractComments(%s, %v) = %v, want %v", code, test.commentType, result, test.expected)
		}
	}
}
