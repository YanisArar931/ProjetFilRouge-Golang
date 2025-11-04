package storage

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
}

type Storer interface {
	AddContact(contact *Contact) error
	GetAllContacts() ([]*Contact, error)
	GetByContactID(id int) (*Contact, error)
	UpdateContact(id int, name, email string) error
	DeleteContact(id int) error
}

var ErrorContactNotFound = func(id int) error {
	return fmt.Errorf("contact avec l'ID %d non trouvé", id)
}