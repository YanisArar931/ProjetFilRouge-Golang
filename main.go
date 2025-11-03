package main

import (
	"fmt"
	"flag"
	"github.com/YanisArar931/ProjetFilRouge-Golang/user"
)

func main() {
	nameFlag := flag.String("name", "", "")
	emailflag := flag.String("email", "", "")
	flag.Parse()

	contactList := user.NewList()

	if *nameFlag != "" && *emailflag != "" {
		newUser := user.User{
			Name: *nameFlag,
			Email: *emailflag,
		}
		contactList.Add(newUser)
		fmt.Printf("Contact créé via les flags : %+v\n", newUser)
		return
	}

	for {
		fmt.Println("-----------------------------")
		fmt.Println("Bienvenue dans notre Mini-CRM !")
		fmt.Println("-----------------------------")
		fmt.Println()
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Supprimer un contact par son ID")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter l'application")

		var choice int
		fmt.Println()
		fmt.Print("Veuillez entrer votre choix (1-5) : ")
		fmt.Println()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			user := user.AddUser()
			contactList.Add(user)
			fmt.Printf("Contact créé : %+v\n", user)
		case 2:
			fmt.Println("Lister tous les contacts sélectionné.")
			contactList.Display()
		case 3:
			contactList.Delete()
		case 4:
			contactList.Update()
		case 5:
			fmt.Println("Merci d'avoir utilisé notre Mini-CRM !")
			return
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}
