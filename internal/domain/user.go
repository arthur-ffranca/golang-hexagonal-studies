package domain

import "errors"

var ErrInvalidUser = errors.New("invalid user")

type User struct {
	ID    string
	Name  string
	Email string
}

func NewUser(id, name, email string) (User, error) {
	if id == "" || name == "" || email == "" {
		return User{}, ErrInvalidUser
	}

	return User{
		ID:    id,
		Name:  name,
		Email: email,
	}, nil
}
