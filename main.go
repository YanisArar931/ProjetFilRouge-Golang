package main

import (
	"flag"
	"fmt"

	contact "github.com/YanisArar931/ProjetFilRouge-Golang/user"
)

func main() {
	contactList := &contact.ContactList{
		Contacts: make(map[int]*contact.Contact),
		NextID:   1,
	}

	nameFlag := flag.String("name", "", "")
	emailFlag := flag.String("email", "", "")
	flag.Parse()

	if *nameFlag != "" && *emailFlag != "" {
		newContact, err := contact.NewContact(contactList.NextID, *nameFlag, *emailFlag)
    if err != nil {
        fmt.Println("Erreur :", err)
    } else {
        contactList.Contacts[contactList.NextID] = newContact
        contactList.NextID++
        fmt.Printf("Contact créé via les flags : [%d] %s <%s>\n", newContact.ID, newContact.Name, newContact.Email)
    }
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
			contactList.AddContact()
		case 2:
			fmt.Println("Lister tous les contacts sélectionné.")
			contactList.DisplayContact()
		case 3:
			contactList.DeleteContact()
		case 4:
			contactList.UpdateContact()
		case 5:
			fmt.Println("Merci d'avoir utilisé notre Mini-CRM !")
			return
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}
