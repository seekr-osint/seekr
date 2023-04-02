package api

func ParsePerson(newPerson Person) Person {
	newPerson = ReplaceNil(newPerson)
	newPerson = CheckMail(newPerson)
	return newPerson
}

func ReplaceNil(newPerson Person) Person {
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
