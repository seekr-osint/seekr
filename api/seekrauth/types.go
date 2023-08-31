package seekrauth

type User struct {
	Username string
	Password string
}
type Users []User

func (u Users) ToMap() map[string]string {
	res := map[string]string{}
	for _, user := range u {
		res[user.Username] = user.Password
	}
	return res
}
