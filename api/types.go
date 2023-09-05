package api

// POST data for the DetectLanguage api call
type LanguageTextInput struct {
				Text string `json:"text" tstype:"string" example:"Hello world"`
}
