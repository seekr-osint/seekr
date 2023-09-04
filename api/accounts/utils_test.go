package accounts

import (
	"fmt"
	"testing"

	"github.com/seekr-osint/seekr/api/tcmultiarg"
)

func TestSetProtocolURL(t *testing.T) {
	testCases := [][2]tcmultiarg.Args{
		[2]tcmultiarg.Args{tcmultiarg.Args{"http://example.com", "https"}, tcmultiarg.Args{"https://example.com", nil}},
		[2]tcmultiarg.Args{tcmultiarg.Args{"https://example.com", ""}, tcmultiarg.Args{"https://example.com", nil}},
		[2]tcmultiarg.Args{tcmultiarg.Args{"example.com", ""}, tcmultiarg.Args{"https://example.com", nil}},
		[2]tcmultiarg.Args{tcmultiarg.Args{"invalid_url", ""}, tcmultiarg.Args{"https://invalid_url", nil}},
		[2]tcmultiarg.Args{tcmultiarg.Args{"", ""}, tcmultiarg.Args{"", fmt.Errorf("invalid url")}},
	}
	tests := tcmultiarg.NewTest(SetProtocolURL, testCases)
	tests.Run(t)
}

func TestParseCheckResult(t *testing.T) {
	testCases := [][2]tcmultiarg.Args{
		[2]tcmultiarg.Args{tcmultiarg.Args{"true"}, tcmultiarg.Args{true, false, nil}},
		[2]tcmultiarg.Args{tcmultiarg.Args{"false"}, tcmultiarg.Args{false, false, nil}},
		[2]tcmultiarg.Args{tcmultiarg.Args{"error: very useful msg"}, tcmultiarg.Args{false, false, fmt.Errorf("very useful msg")}},
		[2]tcmultiarg.Args{tcmultiarg.Args{"rate limited"}, tcmultiarg.Args{false, true, fmt.Errorf("rate limited")}},
		[2]tcmultiarg.Args{tcmultiarg.Args{"kinda true"}, tcmultiarg.Args{false, false, fmt.Errorf("unknown check result: %s", "kinda true")}},
		[2]tcmultiarg.Args{tcmultiarg.Args{""}, tcmultiarg.Args{false, false, fmt.Errorf("empty result")}},
	}
	tests := tcmultiarg.NewTest(ParseCheckResult, testCases)
	tests.Run(t)
}
