package api

func ParsePerson(person Person, config ApiConfig) Person {
	person = person.ReplaceNil()
	person = CheckMail(person, config)
	return person
}

func (person Person) ReplaceNil() Person {
	if person.Pictures == nil {
		person.Pictures = Pictures{}
	}
	if person.Accounts == nil {
		person.Accounts = Accounts{}
	}
	if person.Tags == nil {
		person.Tags = Tags{}
	}
	if person.Sources == nil {
		person.Sources = Sources{}
	}
	if person.Relations == nil {
		person.Relations = Relation{}
	}
	return person
}

func CheckValid(person Person, config ApiConfig) (bool, string) {
	if !person.Civilstatus.IsValid() {
		return false, "civil staus invalid"
	}
	if !person.Religion.IsValid() {
		return false, "invalid religion"
	}
	if !person.SSN.IsValid() {
		return false, "invalid SSN"
	}
	if person.ID == "" {
		return false, "missing id"
	}
	if !person.Gender.IsValid() {
		return false, "invalid gender"
	}
	return true, ""
}
