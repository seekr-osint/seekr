package ethnicity

import (
	"testing"

	"github.com/seekr-osint/seekr/api/enum"
	"github.com/seekr-osint/seekr/api/tc"
)

func isValid(e Ethnicity) bool {
	return e.IsValid()
}

func TestEnum(t *testing.T) {
	test1 := tc.NewEnumIsValidTest(isValid, false, Ethnicity("invalid"), true, African, Asian, CaucasianWhite, HispanicLatino, IndigenousNativeAmerican, MultiracialMixed, NoEthnicity)
	test1.TcTestHandler(t)

	test2 := enum.TcIsValidTest(Enum)
	test2.TcTestHandler(t)
}
