package root

// User entity
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserService interface
type UserService interface {
	CreateUser(u *User) error
	GetByUsername(username string) (*User, error)
}
