package storage

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore(dbPath string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Contact{})
	if err != nil {
		return nil, err
	}

	return &GORMStore{db: db}, nil
}

func (gs *GORMStore) Add(contact *Contact) error {
	result := gs.db.Create(contact)
	return result.Error
}

func (gs *GORMStore) GetAll() ([]*Contact, error) {
	var contacts []*Contact
	result := gs.db.Find(&contacts)
	if result.Error != nil {
		return nil, result.Error
	}
	return contacts, nil
}

func (gs *GORMStore) GetByID(id int) (*Contact, error) {
	var contact Contact
	result := gs.db.First(&contact, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrContactNotFound(id)
		}
		return nil, result.Error
	}
	return &contact, nil
}

func (gs *GORMStore) Update(id int, newName, newEmail string) error {
	_, err := gs.GetByID(id)
	if err != nil {
		return err
	}

	updateData := make(map[string]interface{})
	if newName != "" {
		updateData["name"] = newName
	}
	if newEmail != "" {
		updateData["email"] = newEmail
	}

	if len(updateData) == 0 {
		return nil
	}

	result := gs.db.Model(&Contact{}).Where("id = ?", id).Updates(updateData)
	return result.Error
}

func (gs *GORMStore) Delete(id int) error {
	_, err := gs.GetByID(id)
	if err != nil {
		return err
	}

	result := gs.db.Delete(&Contact{}, id)
	return result.Error
}
