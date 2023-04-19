package ssn

type SSNs map[string]SSN
type SSN struct {
	SSN          string `json:"ssn"`
	State        string `json:"state"`
	areaNumber   int
	groupNumber  int
	serialNumber int
}
