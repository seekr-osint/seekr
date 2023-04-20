package hobby

import "github.com/seekr-osint/seekr/api/functions"

func (hobbies Hobbies) Parse() (Hobbies, error) {
	newHobbies,err := functions.FullParseMapRet(hobbies,"Hobby")
	return newHobbies,err
}

func (hobby Hobby) Parse() (Hobby, error) {
	return hobby, nil
}
