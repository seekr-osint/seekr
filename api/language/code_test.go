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

func TestDetectProgrammingLanguage(t *testing.T) {
	t.Run("Detects Go language", func(t *testing.T) {
		code := `
            package main

            import "fmt"

            func main() {
                fmt.Println("Hello, World!")
            }
        `
		expected := "Go"
		result := DetectProgrammingLanguage(code)
		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})

	t.Run("Detects Python language", func(t *testing.T) {
		code := `
            import numpy as np

            def calculate_mean(numbers):
                return np.mean(numbers)
        `
		expected := "Python"
		result := DetectProgrammingLanguage(code)
		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})

	t.Run("Detects JavaScript language", func(t *testing.T) {
		code := `
						import { greet2 } from './js-sucks.js'
            const greet = () => {
                console.log("Hello, World!");
            }
            
            greet();
        `
		expected := "JavaScript"
		result := DetectProgrammingLanguage(code)
		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})

	t.Run("Detects unknown language", func(t *testing.T) {
		code := `
            // Some code that doesn't match any language patterns
            const x = 5;
        `
		expected := "Unknown"
		result := DetectProgrammingLanguage(code)
		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})
}
