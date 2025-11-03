package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	ID        int
	Name 	  string
	Email     string
}

var nextID = 1

func AddUser() User {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Entrez le nom : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Entrez l'email : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	user := User{
		ID:    nextID,
		Name:  name,
		Email: email,
	}
	nextID++

	fmt.Printf("Contact ajouté: %s %s\n", user.Name, user.Email)
	return user
}
