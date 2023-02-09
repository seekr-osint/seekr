package api

import (
	"context"
	"github.com/rocketlaunchr/google-search"
)

type Search struct {
	Url         string `json:"url"`
	Description string `json:"description"`
	Title       string `json:"title"`
}
type Searches []Search

func SearchString(query string) Searches {
	ctx := context.Background()
	results, _ := googlesearch.Search(ctx, query)
	searches := Searches{}
	for i := 0; i < len(results); i++ {
		searches = append(searches, Search{
			Url:         results[i].URL,
			Description: results[i].Description,
			Title:       results[i].Title,
		})
	}
	return searches

}
