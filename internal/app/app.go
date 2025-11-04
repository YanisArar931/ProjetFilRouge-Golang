package app

import (
	"bufio"
	//"flag"
	"fmt"
	"os"
	"strings"

	// "github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage"
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage"
	//contact "github.com/YanisArar931/ProjetFilRouge-Golang/user"
)

func Run() {
	store := storage.NewMemoryStore()
	// contactList := &contact.ContactList{
	// 	Contacts: make(map[int]*contact.Contact),
	// 	NextID:   1,
	// }

	// nameFlag := flag.String("name", "", "")
	// emailFlag := flag.String("email", "", "")
	// flag.Parse()

	// if *nameFlag != "" && *emailFlag != "" {
	// 	newContact, err := contact.NewContact(contactList.NextID, *nameFlag, *emailFlag)
	// 	if err != nil {
	// 		fmt.Println("Erreur :", err)
	// 	} else {
	// 		contactList.Contacts[contactList.NextID] = newContact
	// 		contactList.NextID++
	// 		fmt.Printf("Contact créé via les flags : [%d] %s <%s>\n", newContact.ID, newContact.Name, newContact.Email)
	// 	}
	// }

	for {
		fmt.Println("-----------------------------")
		fmt.Println("Bienvenue dans notre Mini-CRM !")
		fmt.Println("-----------------------------")
		fmt.Println()
		fmt.Println("1. Ajouter un contact")
		fmt.Println("2. Lister tous les contacts")
		fmt.Println("3. Supprimer un contact par son ID")
		fmt.Println("4. Lister un contact par ID")
		fmt.Println("5. Mettre à jour un contact")
		fmt.Println("6. Quitter l'application")

		var choice int
		fmt.Println()
		fmt.Print("Veuillez entrer votre choix (1-5) : ")
		fmt.Println()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			handleAddContact(store)
		case 2:
			fmt.Println("Lister tous les contacts sélectionné.")
			handleGetAllContacts(store)
		case 3:
			handleDeleteContact(store)
		case 4:
			handleGetByContactID(store)
		case 5:
			handleUpdateContact(store)
		case 6:
			fmt.Println("Merci d'avoir utilisé notre Mini-CRM !")
			return
		default:
			fmt.Println("Choix invalide. Veuillez réessayer.")
		}
	}
}

func handleAddContact(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Entrez le nom : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Entrez l'email : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	newContact := &storage.Contact{
		Name:  name,
		Email: email,
	}

	err := store.AddContact(newContact)
	if err != nil {
		fmt.Println("Erreur lors de l'ajout du contact :", err)
		return
	}

	fmt.Printf("Nouveau contact ajouté : [%d] %s <%s>\n", newContact.ID, newContact.Name, newContact.Email)
}

func handleGetAllContacts(store storage.Storer) {
	contacts, err := store.GetAllContacts()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des contacts :", err)
		return
	}

	fmt.Println("Liste des contacts :")
	for _, contact := range contacts {
		fmt.Printf("- id : %d, Nom : %s, Email : %s\n", contact.ID, contact.Name, contact.Email)
	}
}

func handleGetByContactID(store storage.Storer) {
	fmt.Print("Entrez l'ID du contact à rechercher : ")
	var id int
	fmt.Scan(&id)
	contact, err := store.GetByContactID(id)
	if err != nil {
		fmt.Println("Erreur lors de la récupération du contact :", err)
		return
	}
	fmt.Printf("Contact trouvé : [%d] %s <%s>\n", contact.ID, contact.Name, contact.Email)
}

func handleUpdateContact(store storage.Storer) {
	fmt.Print("Entrez l'ID du contact à mettre à jour : ")
	var id int
	fmt.Scan(&id)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez le nouveau nom : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Entrez le nouvel email : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	err := store.UpdateContact(id, name, email)
	if err != nil {
		fmt.Println("Erreur lors de la mise à jour du contact :", err)
		return
	}
	fmt.Printf("Contact avec l'ID %d mis à jour avec succès.\n", id)
}

func handleDeleteContact(store storage.Storer) {
	fmt.Print("Entrez l'ID du contact à supprimer : ")
	var id int
	fmt.Scan(&id)
	err := store.DeleteContact(id)
	if err != nil {
		fmt.Println("Erreur lors de la récupération du contact :", err)
		return
	}
	fmt.Printf("Contact avec l'ID %d supprimé avec succès.\n", id)
}
