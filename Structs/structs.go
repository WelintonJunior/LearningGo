package main

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

func (u User) outputUserDetailes() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(userFirstName, userLastName, userBirthDate string) (*User, error) {

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

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birth date (DD/MM/YYYY): ")

	var appUser *User

	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		panic(err)
	}

	appUser.outputUserDetailes()
	appUser.clearUserName()
	appUser.outputUserDetailes()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
