package gender

import (
	"testing"

	"github.com/seekr-osint/seekr/api/enum"
	"github.com/seekr-osint/seekr/api/tc"
)

func isValid(g Gender) bool {
	return g.IsValid()
}

func TestEnum(t *testing.T) {
	test1 := tc.NewEnumIsValidTest(isValid, false, Gender("invalid"), true, Male, Female, OtherGender, NoGender)
	test1.TcTestHandler(t)

	test2 := enum.TcIsValidTest(Enum)
	test2.TcTestHandler(t)
}
