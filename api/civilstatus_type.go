package api

type CivilStatus string

const (
	Single    CivilStatus = "Single"
	Married   CivilStatus = "Married"
	Widowed   CivilStatus = "Widowed"
	Divorced  CivilStatus = "Divorced"
	Separated CivilStatus = "Separated"
  None CivilStatus = ""
)

func (cs CivilStatus) IsValid() bool {
	switch cs {
	case Single, Married, Widowed, Divorced, Separated, None:
		return true
	}
	return false
}
