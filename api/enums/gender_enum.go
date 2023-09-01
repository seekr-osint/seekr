package enums

type Gender string

func (gender Gender) Values() []Gender {
	return []Gender{"Male", "Female", "Other"}
}

func (gender Gender) NullValue() Gender {
	return ""
}
