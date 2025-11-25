package main

import (
	"log"

	"github.com/YanisArar931/ProjetFilRouge-Golang/cmd"
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage"
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage/gormstore"
)

type DualStore struct {
	jsonStore storage.Storer
	gormStore storage.Storer
}

func (d *DualStore) AddContact(c *storage.Contact) error {
	if err := d.jsonStore.AddContact(c); err != nil {
		return err
	}
	if err := d.gormStore.AddContact(c); err != nil {
		return err
	}
	return nil
}

func (d *DualStore) GetAllContacts() ([]*storage.Contact, error) {
	return d.jsonStore.GetAllContacts()
}

func (d *DualStore) GetByContactID(id int) (*storage.Contact, error) {
	return d.jsonStore.GetByContactID(id)
}

func (d *DualStore) UpdateContact(id int, name, email string) error {
	if err := d.jsonStore.UpdateContact(id, name, email); err != nil {
		return err
	}
	return d.gormStore.UpdateContact(id, name, email)
}

func (d *DualStore) DeleteContact(id int) error {
	if err := d.jsonStore.DeleteContact(id); err != nil {
		return err
	}
	return d.gormStore.DeleteContact(id)
}

func main() {
	jsonStore, err := storage.NewJSONStorage("contacts.json")
	if err != nil {
		log.Fatal(err)
	}

	gormStore, err := gormstore.NewStore("contacts.db")
	if err != nil {
		log.Fatal(err)
	}

	dual := &DualStore{
		jsonStore: jsonStore,
		gormStore: gormStore,
	}

	cmd.SetStorage(dual)
	cmd.Execute()
}
