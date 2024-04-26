package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u User) OutputUserDetailes() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func New(userFirstName, userLastName, userBirthDate string) (*User, error) {

	if userFirstName == "" || userLastName == "" || userBirthDate == "" {
		return nil, errors.New("Invalid Inputs")
	}

	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}, nil
}
