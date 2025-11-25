package gormstore

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage"
)

type Store struct {
	db *gorm.DB
}

func NewStore(dbPath string) (*Store, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&storage.Contact{}); err != nil {
		return nil, err
	}

	return &Store{db: db}, nil
}

// Implémente storage.Storer

func (s *Store) AddContact(c *storage.Contact) error {
	return s.db.Create(c).Error
}

func (s *Store) GetAllContacts() ([]*storage.Contact, error) {
	var contacts []*storage.Contact
	err := s.db.Find(&contacts).Error
	return contacts, err
}

func (s *Store) GetByContactID(id int) (*storage.Contact, error) {
	var c storage.Contact
	if err := s.db.First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storage.ErrorContactNotFound(id)
		}
		return nil, err
	}
	return &c, nil
}

func (s *Store) UpdateContact(id int, name, email string) error {
	var c storage.Contact
	if err := s.db.First(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return storage.ErrorContactNotFound(id)
		}
		return err
	}
	c.Name = name
	c.Email = email
	return s.db.Save(&c).Error
}

func (s *Store) DeleteContact(id int) error {
	result := s.db.Delete(&storage.Contact{}, id)
	if result.RowsAffected == 0 {
		return storage.ErrorContactNotFound(id)
	}
	return result.Error
}

func NewGORMStore(db *gorm.DB) *Store {
	return &Store{db: db}
}