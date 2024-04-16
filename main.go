package main

import (
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name       string
	Age        int
	Occupation string
}

func main() {

	content, err := os.ReadFile("SAMPLE.TXT")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")

	var users []User

	for i := 0; i < len(lines); i += 4 {
		name := strings.TrimPrefix(strings.TrimSpace(lines[i]), "Name:")
		age := 0
		fmt.Sscanf(strings.TrimPrefix(strings.TrimSpace(lines[i+1]), "Age:"), "%d", &age)
		occupation := strings.TrimPrefix(strings.TrimSpace(lines[i+2]), "Occupation:")

		user := User{Name: name, Age: age, Occupation: occupation}
		users = append(users, user)
	}

	for _, u := range users {
		fmt.Println("Name:", u.Name)
		fmt.Println("Age:", u.Age)
		fmt.Println("Occupation:", u.Occupation)
		fmt.Println()
	}
}
