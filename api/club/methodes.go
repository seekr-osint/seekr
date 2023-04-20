package club

import "github.com/seekr-osint/seekr/api/functions"

func (clubs Clubs) Parse() (Clubs, error) {
	newHobbies,err := functions.FullParseMapRet(clubs,"club")
	return newHobbies,err
}

func (club Club) Parse() (Club, error) {
	return club, nil
}
