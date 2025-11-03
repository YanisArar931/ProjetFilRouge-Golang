package user

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	ID    int
	Name  string
	Email string
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

func (l *List) Update() {
	if len(l.Users) == 0 {
		fmt.Println("Liste vide, rien à mettre à jour")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez l'ID du contact à modifier : ")
	var id int
	fmt.Scan(&id)

	var found *User
	for i := range l.Users {
		if l.Users[i].ID == id {
			found = &l.Users[i]
			break
		}
	}

	if found == nil {
		fmt.Println("Aucun contact trouvé avec cet ID")
		return
	}

	fmt.Printf("Nom actuel : %s. Nouveau nom (laisser vide pour ne pas changer) : ", found.Name)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name != "" {
		found.Name = name
	}

	fmt.Printf("Email actuel : %s. Nouvel email (laisser vide pour ne pas changer) : ", found.Email)
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		found.Email = email
	}

	fmt.Println("Contact mis à jour !")
}

func (l *List) Delete() {
	var id int
	fmt.Print("Entrez l'ID du contact à supprimer : ")
	fmt.Scan(&id)

	for i, user := range l.Users {
		if user.ID == id {
			l.Users = append(l.Users[:i], l.Users[i+1:]...)
			fmt.Printf("Contact avec ID %d supprimé.\n", id)
			return
		}
	}
	fmt.Printf("Aucun contact trouvé avec l'ID %d.\n", id)
}
