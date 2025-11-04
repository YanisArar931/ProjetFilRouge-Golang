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

type List struct {
	Users map[int]User
	nextID int
}

func NewList() *List {
	return &List{
		Users:  make(map[int]User),
		nextID: 1,
	}
}

func AddUser() User {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Entrez le nom : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Entrez l'email : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	return User{
		Name: name,
		Email: email,
	}
}

func (l *List) Add(u User) {
	u.ID = l.nextID
	l.Users[u.ID] = u
	l.nextID++
	fmt.Printf("Nouveau contact ajouté : [%d] %s <%s>\n", u.ID, u.Name, u.Email)
}

func (l *List) Display() {
	if len(l.Users) == 0 {
		fmt.Println("Liste vide")
		return
	}

	fmt.Println("Liste des contacts :")
	for _, u := range l.Users {
		fmt.Printf("- id : %d, Nom : %s, Email : %s\n", u.ID, u.Name, u.Email)
	}
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

	user, exists := l.Users[id]
	if !exists {
		fmt.Println("Aucun contact trouvé avec cet ID.")
		return
	}

	fmt.Printf("Nom actuel : %s. Nouveau nom (laisser vide pour ne pas changer) : ", user.Name)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name != "" {
		user.Name = name
	}

	fmt.Printf("Email actuel : %s. Nouvel email (laisser vide pour ne pas changer) : ", user.Email)
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		user.Email = email
	}
	l.Users[id] = user
	fmt.Println("Contact mis à jour !")
}

func (l *List) Delete() {
	if len(l.Users) == 0 {
		fmt.Println("Liste vide, rien à supprimer")
		return
	}

	var id int
	fmt.Print("Entrez l'ID du contact à supprimer : ")
	fmt.Scan(&id)

	if _, exists := l.Users[id]; exists {
		delete(l.Users, id)
		fmt.Printf("Contact avec ID %d supprimé.\n", id)
	} else {
		fmt.Printf("Aucun contact trouvé avec l'ID %d.\n", id)
	}
}
