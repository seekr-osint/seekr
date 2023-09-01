package enums

type Ethnicity string

func (ethnicity Ethnicity) Values() []Ethnicity {
	return []Ethnicity{"African", "Asian", "Caucasian/White", "Hispanic/Latino", "Indigenous/Native American", "Multiracial/Mixed"}
}

func (ethnicity Ethnicity) NullValue() Ethnicity {
	return ""
}
