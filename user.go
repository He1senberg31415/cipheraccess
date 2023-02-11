package main

import "fmt"

type User struct {
	name         string
	passwordHash string
}

var users = make(map[string]User)

func printUser(userStruct User) {
	fmt.Printf("username: %v\n", userStruct.name)
	fmt.Printf("passwordHash: %v\n", userStruct.passwordHash)
	fmt.Println()
}

func modifyUser(username string) {

}

func createUser(username string) {
	var user User
	user.name = username
	user.passwordHash = hash([]byte(passwordInput("Password: ")))
	users[username] = user
}

func setUser(username string) {
	if _, ok := users[username]; ok {
		modifyUser(username)
	} else {
		createUser(username)
	}
}

func getUser(username string) {
	if username == "*" {
		fmt.Println()
		for _, value := range users {
			printUser(value)
		}
	} else {
		if value, ok := users[username]; ok {
			fmt.Println()
			printUser(value)

		} else {
			fmt.Println("Error")
		}
	}
}
