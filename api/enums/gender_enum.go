package enums

type Gender string
type GenderEnum struct{
	Gender Gender `json:"gender" tstype:"'Male' | 'Female' | 'Other' | ''" example:"Male`
}

func (gender Gender) Values() []Gender {
	return []Gender{ "Male", "Female", "Other" }
}

func (gender Gender) NullValue() Gender {
	return ""
}

