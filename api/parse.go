package api

func ParsePerson(newPerson Person, config ApiConfig) Person {
	newPerson = ReplaceNil(newPerson, config)
	newPerson = CheckMail(newPerson, config)
	return newPerson
}

func ReplaceNil(newPerson Person, config ApiConfig) Person {
	if newPerson.Pictures == nil {
		newPerson.Pictures = Pictures{}
	}
	if newPerson.Accounts == nil {
		newPerson.Accounts = Accounts{}
	}
	if newPerson.Tags == nil {
		newPerson.Tags = Tags{}
	}
	if newPerson.Sources == nil {
		newPerson.Sources = Sources{}
	}
	if newPerson.Relations == nil {
		newPerson.Relations = Relation{}
	}
	return newPerson
}
func CheckValid(newPerson Person, config ApiConfig) (bool, string) {
	if !newPerson.Civilstatus.IsValid() {
		return false, "civil staus invalid"
	}
	if !newPerson.Religion.IsValid() {
		return false, "invalid religion"
	}
	if newPerson.ID == "" {
		return false, "missing id"
	}
	return true, ""
}
