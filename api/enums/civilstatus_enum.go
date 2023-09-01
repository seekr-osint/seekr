package enums

type Civilstatus string

func (civilstatus Civilstatus) Values() []Civilstatus {
	return []Civilstatus{"Single", "Married", "Windowed", "Divorced", "Seperated"}
}

func (civilstatus Civilstatus) NullValue() Civilstatus {
	return ""
}
