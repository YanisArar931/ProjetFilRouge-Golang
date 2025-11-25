package storage

type DualStore struct {
	JSON Storer
	GORM Storer
}

func NewDualStore(json, gorm Storer) Storer {
	return &DualStore{
		JSON: json,
		GORM: gorm,
	}
}

func (d *DualStore) AddContact(c *Contact) error {
	_ = d.JSON.AddContact(c)
	return d.GORM.AddContact(c)
}

func (d *DualStore) GetAllContacts() ([]*Contact, error) {
	return d.JSON.GetAllContacts()
}

func (d *DualStore) GetByContactID(id int) (*Contact, error) {
	return d.JSON.GetByContactID(id)
}

func (d *DualStore) UpdateContact(id int, name, email string) error {
	_ = d.JSON.UpdateContact(id, name, email)
	return d.GORM.UpdateContact(id, name, email)
}

func (d *DualStore) DeleteContact(id int) error {
	_ = d.JSON.DeleteContact(id)
	return d.GORM.DeleteContact(id)
}
