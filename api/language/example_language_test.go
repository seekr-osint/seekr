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
func ExampleDetectLanguage_shortText() {
	langs := language.DetectLanguage("Я йду текста")
	fmt.Printf("%s: %.3f\n", "Russian", langs["Russian"])
	fmt.Printf("%s: %.3f\n", "Belarusian", langs["Belarusian"])
	// Output:
	//
	// Russian: 0.248
	// Belarusian: 0.050

}
