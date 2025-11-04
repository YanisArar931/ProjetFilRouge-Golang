package contact

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	ID    int
	Name  string
	Email string
}

type ContactList struct {
	Contacts map[int]*Contact
	NextID   int
}

func NewContact(id int, name, email string) (*Contact, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" {
		return nil, errors.New("le nom ne peut pas être vide")
	}
	if email == "" || !strings.Contains(email, "@") {
		return nil, errors.New("email invalide")
	}

	return &Contact{
		ID:    id,
		Name:  name,
		Email: email,
	}, nil
}

func (c *ContactList) AddContact() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Entrez le nom : ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Entrez l'email : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	newContact, err := NewContact(c.NextID, name, email)
	if err != nil {
		fmt.Println("Erreur lors de la création du contact :", err)
		return
	}
	c.Contacts[c.NextID] = newContact
	c.NextID++
	fmt.Printf("Nouveau contact ajouté : [%d] %s <%s>\n", newContact.ID, newContact.Name, newContact.Email)
}

func (c *ContactList) DisplayContact() {
	if len(c.Contacts) == 0 {
		fmt.Println("Liste vide")
		return
	}

	fmt.Println("Liste des contacts :")
	for _, contact := range c.Contacts {
		fmt.Printf("- id : %d, Nom : %s, Email : %s\n", contact.ID, contact.Name, contact.Email)
	}
}

func (c *ContactList) UpdateContact() {
	if len(c.Contacts) == 0 {
		fmt.Println("Liste vide, rien à mettre à jour")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez l'ID du contact à modifier : ")
	var id int
	fmt.Scan(&id)

	contact, exists := c.Contacts[id]
	if !exists {
		fmt.Println("Aucun contact trouvé avec cet ID.")
		return
	}

	fmt.Printf("Nom actuel : %s. Nouveau nom (laisser vide pour ne pas changer) : ", contact.Name)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name != "" {
		contact.Name = name
	}

	fmt.Printf("Email actuel : %s. Nouvel email (laisser vide pour ne pas changer) : ", contact.Email)
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		contact.Email = email
	}
	fmt.Println("Contact mis à jour !")
}

func (c * ContactList) DeleteContact() {
	if len(c.Contacts) == 0 {
		fmt.Println("Liste vide, rien à supprimer")
		return
	}

	var id int
	fmt.Print("Entrez l'ID du contact à supprimer : ")
	fmt.Scan(&id)

	if _, exists := c.Contacts[id]; exists {
		delete(c.Contacts, id)
		fmt.Printf("Contact avec ID %d supprimé.\n", id)
	} else {
		fmt.Printf("Aucun contact trouvé avec l'ID %d.\n", id)
	}
}
