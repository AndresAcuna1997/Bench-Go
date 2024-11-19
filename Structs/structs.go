package main

import (
	"fmt"

	"examplo.com/structs/user"
)

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userfirstName, userlastName, userbirthdate)

	admin := user.NewAdmin("abc@abc.com", "Lalala")

	if err != nil {
		fmt.Println(err)
		return
	}

	// appUser := User{
	// 	firstName: userfirstName,
	// 	lastName:  userlastName,
	// 	birthdate: userbirthdate,
	// 	createdAt: time.Now(),
	// }

	// appUser = User{
	// 	userfirstName,
	// 	userlastName,
	// 	userbirthdate,
	// }

	admin.OutPutUserDetails()
	appUser.OutPutUserDetails()
	appUser.ClearName()
	appUser.OutPutUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
