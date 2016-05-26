package main

type User struct {
	Username string
	Password string
}

func (u *User) Login() (token string, err error) {
	// ...
	return
}

func (u *User) Logout() error {
	// ...
	return
}
