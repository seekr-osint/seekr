package sources

// Types
type Sources map[string]Source
type Source struct {
	Url string `json:"url"`
}
