package main

import "fmt"

func main(){
	for{
		fmt.Println("-----------------------------")
		fmt.Println("Bienvenue dans notre Mini-CRM !")
		fmt.Println("-----------------------------")
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Supprimer un contact par son ID")
		fmt.Println("4. Mettre à jour un contact")
		fmt.Println("5. Quitter l'application")

		var choice int
		fmt.Print("Veuillez entrer votre choix (1-5) : ")
		fmt.Scan(&choice)

		switch choice {
		case 1: 
			fmt.Println("Ajouter un contact sélectionné.")
		case 2: 
			fmt.Println("Lister tous les contacts sélectionné.")
		case 3: 
			fmt.Println("Supprimer un contact par son ID sélectionné.")
		case 4: 
			fmt.Println("Mettre à jour un contact sélectionné.")
		case 5: 
			fmt.Println("Merci d'avoir utilisé notre Mini-CRM !")
			return
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}