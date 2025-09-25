package storage

import "fmt"

type Contact struct {
	ID    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null"`
}

type Storer interface {
	Add(contact *Contact) error
	GetAll() ([]*Contact, error)
	GetByID(id int) (*Contact, error)
	Update(id int, newName, newEmail string) error
	Delete(id int) error
}

var ErrContactNotFound = func(id int) error {
	return fmt.Errorf("Contact avec l'ID %d non trouvé", id)
}
