package api

import (
	whois "github.com/likexian/whois"
)

func Whois(url string, config ApiConfig) string {
	result, err := whois.Whois(url)
	CheckAndLog(err, "error whois lookup", config)
	return result
}
