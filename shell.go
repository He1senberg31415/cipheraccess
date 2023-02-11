package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return input
}

func passwordInput(prompt string) string {
	fmt.Print(prompt)
	bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error reading password:", err)
	}
	fmt.Println()
	return string(bytePassword)
}

func handleSetCommands(inputSlice []string) {
	switch inputSlice[1] {
	case "user":
		setUser(inputSlice[2])
	case "port":
	case "whitelist":
	case "blacklist":
	case "block":
	}
}

func handleGetCommands(inputSlice []string) {
	switch inputSlice[1] {
	case "user":
		getUser(inputSlice[2])
	case "port":
	case "whitelist":
	case "blacklist":
	case "block":
	}
}

func handleRemoveCommands(inputSlice []string) {

}

func startConfig() {
	for {
		input := input("> ")
		inputSlice := strings.Split(strings.ReplaceAll(input, "\n", ""), " ")
		switch inputSlice[0] {
		case "set":
			handleSetCommands(inputSlice)
		case "get":
			handleGetCommands(inputSlice)
		case "remove":
			handleRemoveCommands(inputSlice)
		case "exit":
			return
		default:
			fmt.Println("This is not a command")
		}
	}
}

func startShell() {
	switch os.Args[1] {
	case "listen":
		// listen for incoming connections
	case "connect":
		// connect to machine
	case "config":
		fmt.Println("To enter the config mode you have to input the master password.")
		_ = passwordInput("Password: ")
		startConfig()
	}
}
