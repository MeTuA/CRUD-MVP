package users

type User struct {
	ID       int    `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name"`
	Phone    string `json:"phone" db:"phone"`
	Address  string `json:"address" db:"address"`
}
