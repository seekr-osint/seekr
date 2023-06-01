package language

import (
	"github.com/pemistahl/lingua-go"
)

func DetectLanguage(text string) map[string]float64 {

	detector := lingua.NewLanguageDetectorBuilder().FromAllLanguages().WithPreloadedLanguageModels().Build()

	confidenceValues := detector.ComputeLanguageConfidenceValues(text)

	confidenceMap := make(map[string]float64)

	for _, elem := range confidenceValues {
		confidenceMap[elem.Language().String()] = elem.Value()
	}
	return confidenceMap
}
