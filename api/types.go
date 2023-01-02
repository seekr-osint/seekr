package api

// main data set
type person struct {
	ID             string             `json:"id"`
	Maidenname     string             `json:"maidenname"`
	Name           string             `json:"name"`
	Age            int                `json:"age"`
	Birthday       string             `json:"bday"`
	Address        string             `json:"address"`
	Phone          string             `json:"phone"`
	SSN            string             `json:"ssn"`
	Civilstatus    string             `json:"civilstatus"`
	Kids           string             `json:"kids"`
	Hobbies        string             `json:"hobbies"`
	Email          string             `json:"email"`
	Occupation     string             `json:"occupation"`
	Prevoccupation string             `json:"prevoccupation"`
	Education      string             `json:"education"`
	Military       string             `json:"military"`
	Religion       string             `json:"religion"`
	Pets           string             `json:"pets"`
	Club           string             `json:"club"`
	Legal          string             `json:"legal"`
	Political      string             `json:"political"`
	SSN            string             `json:"SSN"`
	Education      string             `json:"education"`
	Religion       string             `json:"religion"`
	Pets           string             `json:"pets"`
	MaidenName     string             `json:"maidenName"`
	Notes          string             `json:"notes"`
	Accounts       map[string]Account `json:"accounts"`
	NotAccounts    map[string]Account `json:"notaccounts"`
}

type DataBase map[string]person
