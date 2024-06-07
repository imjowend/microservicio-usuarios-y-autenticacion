package user

type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

type InMemDB map[string]*User
