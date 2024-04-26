package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birth date (DD/MM/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		panic(err)
	}

	admin := user.NewAdmin("Yasmin", "2206")

	admin.User.OutputUserDetailes()
	admin.User.ClearUserName()
	admin.User.OutputUserDetailes()

	appUser.OutputUserDetailes()
	appUser.ClearUserName()
	appUser.OutputUserDetailes()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
