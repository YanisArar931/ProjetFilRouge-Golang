package storage

// import (
// 	"fmt"
// )

type MemoryStore struct {
	contacts map[int]*Contact
	nextId   int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*Contact),
		nextId:   1,
	}

}

func (ms *MemoryStore) AddContact(contact *Contact) error {
	contact.ID = ms.nextId
	ms.contacts[ms.nextId] = contact
	ms.nextId++
	return nil
}

// func (ms *MemoryStore) GetAllContacts(contact *Contact) error {

// }

// func (ms *MemoryStore) GetByContactID(contact *Contact) error {

// }

// func (ms *MemoryStore) UpdateContact(contact *Contact) error {

// }

// func (ms *MemoryStore) DeleteContact(contact *Contact) error {

// }
