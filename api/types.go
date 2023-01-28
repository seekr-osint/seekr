package api

// main data set
type person struct {
	ID             string             `json:"id"`
	Name           string             `json:"name"`
	Pictures       []string           `json:"pictures"`
	Maidenname     string             `json:"maidenname"`
	Age            int                `json:"age"`
	Birthday       string             `json:"bday"`
	Address        string             `json:"address"`
	Phone          string             `json:"phone"`
	SSN            string             `json:"ssn"`
	Civilstatus    string             `json:"civilstatus"`
	Kids           string             `json:"kids"`
	Hobbies        string             `json:"hobbies"`
	Email          Email              `json:"email"`
	Occupation     string             `json:"occupation"`
	Prevoccupation string             `json:"prevoccupation"`
	Education      string             `json:"education"`
	Military       string             `json:"military"`
	Religion       string             `json:"religion"`
	Pets           string             `json:"pets"`
	Club           string             `json:"club"`
	Legal          string             `json:"legal"`
	Political      string             `json:"political"`
	Notes          string             `json:"notes"`
	Relations      Relation           `json:"relations"`
	Sources        []string           `json:"sources"`
	Accounts       Accounts           `json:"accounts"`
	NotAccounts    map[string]Account `json:"notaccounts"`
}

type DataBase map[string]person
type Relation map[string][]string

type Email []struct {
	Mail       string   `json:"mail"`
	Value      int      `json:"value"`
	Src        string   `json:"src"`
	Services   []string `json:"services"`
	Valid      bool     `json:"valid"`
	Gmail      bool     `json:"gmail"`
	ValidGmail bool     `json:"validGmail"`
}

// type Accounts map[string]Account
type Accounts []Account
type Account struct {
	Service   string   `json:"service"`  // example: GitHub
	Id        string   `json:"id"`       // example: 1224234
	Username  string   `json:"username"` // example: 9glenda
	Url       string   `json:"url"`      // example: https://github.com/9glenda
	Picture   []string `json:"profilePicture"`
	ImgHash   []uint64 `json:"imgHash"`
	Bio       []string `json:"bio"`       // example: pro hacka
	Firstname string   `json:"firstname"` // example: Glenda
	Lastname  string   `json:"lastname"`  // example: Belov
	Location  string   `json:"location"`  // example: Moscow
	Created   string   `json:"created"`   // example: 2020-07-31T13:04:48Z
	Updated   string   `json:"updated"`
	Blog      string   `json:"blog"`
	Followers int      `json:"followers"`
	Following int      `json:"following"`
}
