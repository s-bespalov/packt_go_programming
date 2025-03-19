package main

import (
	"fmt"
	"os"
)

func getUsers() map[string]string {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracey",
	}
	return users
}

func getUser(id string) (string, bool) {
	users := getUsers()
	user, exist := users[id]
	return user, exist
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}
	userId := os.Args[1]
	name, exist := getUser(userId)
	if !exist {
		fmt.Printf("Passed user ID (%v) not found\n", userId)
		for key, value := range getUsers() {
			fmt.Println("  ID: ", key, "Name:", value)
		}
		os.Exit(1)
	}
	fmt.Printf("Hi, %v\n", name)
}
