package language_test

import (
	"fmt"

	"github.com/seekr-osint/seekr/api/language"
)

// Taking in some Russian text as input.
func ExampleDetectLanguage() {
	langs := language.DetectLanguage("Нет ни одного пользователя-искателя. Все ненавидят Seekr, даже разработчики его не используют, потому что это просто отстой. API - это просто глючный беспорядок.")
	for name, value := range langs {
		// only printing the detected langauges
		if value > 0 {
			fmt.Printf("%s: %v\n", name, value)
		}
	}
	// Output:
	//
	// Russian: 1
}

// If the text is short the result of the scann will often include multiple langauges. For example this Russian Text can be confued with Belarusian text.
//
// Altho this still is clearly russian text due to the much higher valuse for Russian.
//
// NOTE the e-05 in the Belarusiann text.
func ExampleDetectLanguage_shortText() {
	langs := language.DetectLanguage("часть русского текста также можно спутать с белорусским текстом")
	fmt.Printf("%s: %v\n", "Russian", langs["Russian"])
	fmt.Printf("%s: %v\n", "Belarusian", langs["Belarusian"])
	// Output:
	//
	// Russian: 0.9866925418362897
	// Belarusian: 4.6738631429e-05

}
