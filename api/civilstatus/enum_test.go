package civilstatus

import (
	"testing"

	"github.com/seekr-osint/seekr/api/enum"
)

func TestEnum(t *testing.T) {
	test := enum.TcIsValidTest(Enum)
	test.TcTestHandler(t)
}
