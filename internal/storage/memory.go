package storage

import "fmt"

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

func (ms *MemoryStore) GetAllContacts() ([]*Contact, error) {
	if len(ms.contacts) == 0 {
		return nil, fmt.Errorf("Aucun contact trouvé")
	}
	contacts := []*Contact{}
	for _, c := range ms.contacts {
		contacts = append(contacts, c)
	}
	return contacts, nil
}

func (ms *MemoryStore) GetByContactID(id int) (*Contact, error) {
	contact, exists := ms.contacts[id]
	if !exists {
		return nil, fmt.Errorf("contact avec l'ID %d non trouvé", id)
	}
	return contact, nil
}

func (ms *MemoryStore) UpdateContact(id int, name, email string) error {
	contact, exists := ms.contacts[id]
	if !exists {
		return fmt.Errorf("contact avec l'ID %d non trouvé", contact.ID)
	}
	contact.Name = name
	contact.Email = email
	return nil
}

func (ms *MemoryStore) DeleteContact(id int) error {
	contact, exists := ms.contacts[id]
	if !exists {
		return fmt.Errorf("contact avec l'ID %d non trouvé", contact.ID)
	}
	delete(ms.contacts, contact.ID)
	return nil
}
