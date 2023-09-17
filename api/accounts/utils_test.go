package accounts

import (
	"fmt"
	"testing"

	"github.com/seekr-osint/seekr/api/tcmultiarg"
)

func TestSetProtocolURL(t *testing.T) {
	testCases := [][2]tcmultiarg.Args{
		{{"http://example.com", "https"}, tcmultiarg.Args{"https://example.com", nil}},
		{{"https://example.com", ""}, tcmultiarg.Args{"https://example.com", nil}},
		{{"example.com", ""}, tcmultiarg.Args{"https://example.com", nil}},
		{{"invalid_url", ""}, tcmultiarg.Args{"https://invalid_url", nil}},
		{{"", ""}, tcmultiarg.Args{"", fmt.Errorf("invalid url")}},
	}
	tests := tcmultiarg.NewTest(SetProtocolURL, testCases)
	tests.Run(t)
}

func TestParseCheckResult(t *testing.T) {
	testCases := [][2]tcmultiarg.Args{
		{{"true"}, tcmultiarg.Args{true, false, nil}},
		{{"false"}, tcmultiarg.Args{false, false, nil}},
		{{"error: very useful msg"}, tcmultiarg.Args{false, false, fmt.Errorf("very useful msg")}},
		{{"error: "}, tcmultiarg.Args{false, false, fmt.Errorf("error message missing")}},
		{{"rate limited"}, tcmultiarg.Args{false, true, fmt.Errorf("rate limited")}},
		{{"kinda true"}, tcmultiarg.Args{false, false, fmt.Errorf("unknown check result: %s", "kinda true")}},
		{{""}, tcmultiarg.Args{false, false, fmt.Errorf("empty result")}},
	}
	tests := tcmultiarg.NewTest(ParseCheckResult, testCases)
	tests.Run(t)
}
