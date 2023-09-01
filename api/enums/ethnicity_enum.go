package enums

type Ethnicity string
type EthnicityEnum struct{
	Ethnicity Ethnicity `json:"ethnicity" tstype:"'African' | 'Asian' | 'Caucasian/White' | 'Hispanic/Latino' | 'Indigenous/Native American' | 'Multiracial/Mixed' | ''" example:"African`
}

func (ethnicity Ethnicity) Values() []Ethnicity {
	return []Ethnicity{ "African", "Asian", "Caucasian/White", "Hispanic/Latino", "Indigenous/Native American", "Multiracial/Mixed" }
}

func (ethnicity Ethnicity) NullValue() Ethnicity {
	return ""
}

