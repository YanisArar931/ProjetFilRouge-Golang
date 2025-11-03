package main

import (
	"fmt"

	"github.com/YanisArar931/ProjetFilRouge-Golang/user"
)

func main() {
	contactList := user.List{}

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
			fmt.Printf("Utilisateur créé : %+v\n", user)
		case 2:
			fmt.Println("Lister tous les contacts sélectionné.")
			contactList.Display()
		case 3:
			fmt.Println("Supprimer un contact par son ID sélectionné.")
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
