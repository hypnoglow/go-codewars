package main

import (
	"encoding/json"
	"fmt"

	"github.com/hypnoglow/go-codewars"
)

// Insert your token here.
const token = ""

func main() {
	cw := codewars.NewClient(token)

	usernames := [2]string{"Hypnoglow", "hypnoglow123456"}

	for _, username := range usernames {
		fmt.Printf("Fetching user data for username `%s`...\n", username)
		user, _, err := cw.Users.GetUser(username)

		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
		} else {
			userJSON, _ := json.MarshalIndent(user, "", "  ")
			fmt.Printf("User:\n%s\n\n", userJSON)
		}
	}
}
