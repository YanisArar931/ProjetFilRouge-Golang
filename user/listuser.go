package user

import "fmt"

type List struct {
	Users []User
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

func (l *List) Add(u User) {
	l.Users = append(l.Users, u)
}
