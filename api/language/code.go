package language

import (
	"fmt"
	"regexp"
	"strings"
)

type AnalyzedComment struct {
	Text string             `json:"text"`
	Lang map[string]float64 `json:"lang"`
}
type CommentType struct {
	Regex   string `json:"regex"`
	Replace []string
}
type Lang []CommentType

var Langs map[string][]CommentType = map[string][]CommentType{
	"python": {Hash},
	"go":     {DoubleSlash, DoubleSlashMultiLine},
	"js":     {DoubleSlash, DoubleSlashMultiLine},
}
var (
	DoubleSlash = CommentType{
		Regex:   `(?m)^\s*//.*?$`,
		Replace: []string{"//"},
	}
	DoubleSlashMultiLine = CommentType{
		Regex:   `(?s)/\*.*?\*/`,
		Replace: []string{"/*", "*/"},
	}
	Hash = CommentType{
		Regex:   `(?m)^\s*#.*?$`,
		Replace: []string{"#"},
	}
)

func ExtractComments(code string, commentType ...CommentType) []string {
	comments := []string{}
	for _, pattern := range commentType {
		commentPattern := string(pattern.Regex)

		r := regexp.MustCompile(commentPattern)
		found := r.FindAllString(code, -1)
		for _, comment := range found {
			for _, replacePattern := range pattern.Replace {
				comment = strings.ReplaceAll(comment, replacePattern, "")
			}
			comments = append(comments, comment)
		}

	}
	return comments
}

func AnalyzeCode(code string, lang string) []AnalyzedComment {
	analyzedComments := []AnalyzedComment{}
	comments := ExtractComments(code, Langs[lang]...)
	for _, comment := range comments {
		fmt.Printf("%s\n", comment)
		analyzed := AnalyzedComment{
			Text: comment,
			Lang: DetectLanguage(string(comment)),
		}
		analyzedComments = append(analyzedComments, analyzed)
	}
	return analyzedComments
}
