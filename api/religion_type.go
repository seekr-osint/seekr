package api

import (
  "strings"
  "fmt"
)

type Religion string

const (
	Christianity Religion = "Christianity"
	Islam        Religion = "Islam"
	Hinduism     Religion = "Hinduism"
	Buddhism     Religion = "Buddhism"
	Sikhism      Religion = "Sikhism"
	Judaism      Religion = "Judaism"
	Other        Religion = "Other"
	NoReligion   Religion = ""
	Atheism      Religion = "Atheism"
)

type Sect string

const (
	Catholic     Sect = "Catholic"
	Protestant   Sect = "Protestant"
	Sunni        Sect = "Sunni"
	Shia         Sect = "Shia"
	Theravada    Sect = "Theravada"
	Mahayana     Sect = "Mahayana"
	Vajrayana    Sect = "Vajrayana"
	Amish        Sect = "Amish"
	Reform       Sect = "Reform"
	Conservative Sect = "Conservative"
	Orthodox     Sect = "Orthodox"
	Hasidic      Sect = "Hasidic"
)

var validSects = map[Religion][]Sect{
	Christianity: {Catholic, Protestant},
	Islam:        {Sunni, Shia},
	Hinduism:     {},
	Buddhism:     {Theravada, Mahayana, Vajrayana},
	Sikhism:      {},
	Judaism:      {Amish, Reform, Conservative, Orthodox, Hasidic},
	Other:        {},
	NoReligion:   {},
	Atheism:      {},
}


func (religion Religion) Markdown() string {
	var sb strings.Builder
	if religion.IsValid() && religion != "" {
		sb.WriteString(fmt.Sprintf("- Religion: `%s`\n",religion ))
	}
  return sb.String()
}
func (r Religion) IsValid() bool {
	switch r {
	case Christianity, Islam, Hinduism, Buddhism, Sikhism, Judaism, Other, NoReligion, Atheism:
		return true
	}
	return false
}

func (r Religion) IsValidSect(s Sect) bool {
	if !r.IsValid() {
		return false
	}
	for _, validSect := range validSects[r] {
		if s == validSect {
			return true
		}
	}
	return false
}
