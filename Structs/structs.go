package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	user := new(User)
	user.firstName = getUserData("Please enter your first name: ")
	user.lastName = getUserData("Please enter your last name: ")
	user.birthDate = getUserData("Please enter your birth date (DD/MM/YYYY): ")
	user.createdAt = time.Now()

	outputUserDetailes(user)
}

func outputUserDetailes(u *User) {
	fmt.Println(u.firstName, u.lastName, u.birthDate, u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
