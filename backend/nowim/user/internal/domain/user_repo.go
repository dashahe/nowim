package domain

type UserRepository interface {
	AddUser(user *User) error
	UserByName(username string) (*User, error)
	UserByID(id int64) (*User, error)
	AllUsers() (map[int64]*User, error)
}
