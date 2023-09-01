package enums

type Civilstatus string
type CivilstatusEnum struct{
	Civilstatus Civilstatus `json:"civilstatus" tstype:"'Single' | 'Married' | 'Windowed' | 'Divorced' | 'Seperated' | ''" example:"Single`
}

func (civilstatus Civilstatus) Values() []Civilstatus {
	return []Civilstatus{ "Single", "Married", "Windowed", "Divorced", "Seperated" }
}

func (civilstatus Civilstatus) NullValue() Civilstatus {
	return ""
}

