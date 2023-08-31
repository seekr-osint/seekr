package names

type Name struct {
	Name string `json:"name" genji:"name"`
	ID   uint   `json:"id" genji:"id"`
}

type Names []Name

func (n Names) UnusedID() int {
	bigges := 0
	for _, i := range n {
		if i.ID > uint(bigges) {
			bigges = int(i.ID)
		}
	}
	return bigges + 1
}
