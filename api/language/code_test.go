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
		filename := "main.go"
		expected := "Go"
		result := DetectProgrammingLanguage(code, filename)
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
		filename := "main.py"
		expected := "Python"
		result := DetectProgrammingLanguage(code, filename)
		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})

	t.Run("Detects Jave language without filename", func(t *testing.T) {
		code := `
package tech.sourced.enry;

import org.junit.Test;

import static org.junit.Assert.*;

public class EnryTest {

    @Test
    public void getLanguage() {
        String code = "<?php $foo = bar();";
        assertEquals("PHP", Enry.getLanguage("foobar.php", code.getBytes()));
    }

    @Test
    public void getLanguageWithNullContent() {
        assertEquals("Python", Enry.getLanguage("foo.py",  null));
    }

    @Test
    public void getLanguageWithEmptyContent() {
        assertEquals("Go", Enry.getLanguage("baz.go",  "".getBytes()));
        assertEquals("Go", Enry.getLanguage("baz.go",  null));
    }

    @Test
    public void getLanguageWithNullFilename() {
        byte[] content = "#!/usr/bin/env python".getBytes();
        assertEquals("Python", Enry.getLanguage(null, content));
    }

    @Test
    public void getLanguageByContent() {
        String code = "<?php $foo = bar();";
        assertGuess(
                "PHP",
                true,
                Enry.getLanguageByContent("foo.php", code.getBytes())
        );
    }

    @Test
    public void getLanguageByFilename() {
        assertGuess(
                "Maven POM",
                true,
                Enry.getLanguageByFilename("pom.xml")
        );
    }

    @Test
    public void getLanguageByEmacsModeline() {
        String code = "// -*- font:bar;mode:c++ -*-\n" +
                "template <typename X> class { X i; };";
        assertGuess(
                "C++",
                true,
                Enry.getLanguageByEmacsModeline(code.getBytes())
        );
    }

    @Test
    public void getLanguageByExtension() {
        assertGuess(
                "Ruby",
                true,
                Enry.getLanguageByExtension("foo.rb")
        );
    }

    @Test
    public void getLanguageByShebang() {
        String code = "#!/usr/bin/env python";
        assertGuess(
                "Python",
                true,
                Enry.getLanguageByShebang(code.getBytes())
        );
    }

    @Test
    public void getLanguageByModeline() {
        String code = "// -*- font:bar;mode:c++ -*-\n" +
                "template <typename X> class { X i; };";
        assertGuess(
                "C++",
                true,
                Enry.getLanguageByModeline(code.getBytes())
        );

        code = "# vim: noexpandtab: ft=javascript";
        assertGuess(
                "JavaScript",
                true,
                Enry.getLanguageByModeline(code.getBytes())
        );
    }

    @Test
    public void getLanguageByVimModeline() {
        String code = "# vim: noexpandtab: ft=javascript";
        assertGuess(
                "JavaScript",
                true,
                Enry.getLanguageByVimModeline(code.getBytes())
        );
    }

    @Test
    public void getLanguageExtensions() {
        String[] exts = Enry.getLanguageExtensions("Go");
        String[] expected = {".go"};
        assertArrayEquals(expected, exts);
    }

    @Test
    public void getLanguages() {
        String code = "#include <stdio.h>" +
                "" +
                "extern int foo(void *bar);";

        String[] result = Enry.getLanguages("foo.h", code.getBytes());
        String[] expected = {"C", "C++", "Objective-C"};
        assertArrayEquals(expected, result);
    }

    @Test
    public void getMimeType() {
        assertEquals(
                "text/x-ruby",
                Enry.getMimeType("foo.rb", "Ruby")
        );
    }

    @Test
    public void isBinary() {
        assertFalse(Enry.isBinary("hello = 'world'".getBytes()));
    }

    @Test
    public void isConfiguration() {
        assertTrue(Enry.isConfiguration("config.yml"));
        assertFalse(Enry.isConfiguration("FooServiceProviderImplementorFactory.java"));
    }

    @Test
    public void isDocumentation() {
        assertTrue(Enry.isDocumentation("docs/"));
        assertFalse(Enry.isDocumentation("src/"));
    }

    @Test
    public void isDotFile() {
        assertTrue(Enry.isDotFile(".env"));
        assertFalse(Enry.isDotFile("config.json"));
    }

    @Test
    public void isImage() {
        assertTrue(Enry.isImage("yup.jpg"));
        assertFalse(Enry.isImage("nope.go"));
    }

    @Test
    public void getColor() {
        assertEquals(
                "#00ADD8",
                Enry.getColor("Go")
        );
    }

    void assertGuess(String language, boolean safe, Guess guess) {
        assertEquals(language, guess.language);
        assertEquals(safe, guess.safe);
    }

}
        `
		filename := "main.java"
		expected := "Java"
		result := DetectProgrammingLanguage(code, filename)
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
		filename := "main.js"
		expected := "JavaScript"
		result := DetectProgrammingLanguage(code, filename)
		if result != expected {
			t.Errorf("Expected %s, but got %s", expected, result)
		}
	})
}
