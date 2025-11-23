package storage

import (
    "encoding/json"
    "os"
)

type JSONStorage struct {
    FilePath string
    contacts []*Contact
}

func NewJSONStorage(path string) (*JSONStorage, error) {
    s := &JSONStorage{FilePath: path}

    err := s.load()
    if err != nil {
        return nil, err
    }

    return s, nil
}

func (s *JSONStorage) load() error {
    data, err := os.ReadFile(s.FilePath)
    if err != nil {
        if os.IsNotExist(err) {
            s.contacts = []*Contact{}
            return nil
        }
        return err
    }

    return json.Unmarshal(data, &s.contacts)
}

func (s *JSONStorage) save() error {
    data, err := json.MarshalIndent(s.contacts, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(s.FilePath, data, 0644)
}

func (s *JSONStorage) AddContact(contact *Contact) error {
    if len(s.contacts) > 0 {
        contact.ID = s.contacts[len(s.contacts)-1].ID + 1
    } else {
        contact.ID = 1
    }

    s.contacts = append(s.contacts, contact)
    return s.save()
}

func (s *JSONStorage) GetAllContacts() ([]*Contact, error) {
    return s.contacts, nil
}

func (s *JSONStorage) GetByContactID(id int) (*Contact, error) {
    for _, c := range s.contacts {
        if c.ID == id {
            return c, nil
        }
    }
    return nil, ErrorContactNotFound(id)
}

func (s *JSONStorage) UpdateContact(id int, name string, email string) error {
    for _, c := range s.contacts {
        if c.ID == id {
            c.Name = name
            c.Email = email
            return s.save()
        }
    }
    return ErrorContactNotFound(id)
}

func (s *JSONStorage) DeleteContact(id int) error {
    for i, c := range s.contacts {
        if c.ID == id {
            s.contacts = append(s.contacts[:i], s.contacts[i+1:]...)
            return s.save()
        }
    }
    return ErrorContactNotFound(id)
}
