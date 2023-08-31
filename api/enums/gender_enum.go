package enums

type Gender string

func (gender Gender) Values() []Gender {
	return []Gender{ "male", "female", "other" }
}

func (gender Gender) NullValue() Gender {
	return ""
}

