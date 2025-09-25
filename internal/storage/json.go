package storage

import (
	"encoding/json"
	"errors"
	"os"
)

type JSONStore struct {
	contacts []*Contact
	nextID   int
	filepath string
}

func NewJSONStore(filepath string) (*JSONStore, error) {
	store := &JSONStore{
		contacts: []*Contact{},
		nextID:   1,
		filepath: filepath,
	}

	if err := store.load(); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	for _, c := range store.contacts {
		if c.ID >= store.nextID {
			store.nextID = c.ID + 1
		}
	}

	return store, nil
}

func (js *JSONStore) Add(contact *Contact) error {
	contact.ID = js.nextID
	js.contacts = append(js.contacts, contact)
	js.nextID++
	return js.save()
}

func (js *JSONStore) GetAll() ([]*Contact, error) {
	return js.contacts, nil
}

func (js *JSONStore) GetByID(id int) (*Contact, error) {
	for _, c := range js.contacts {
		if c.ID == id {
			return c, nil
		}
	}
	return nil, ErrContactNotFound(id)
}

func (js *JSONStore) Update(id int, newName, newEmail string) error {
	contact, err := js.GetByID(id)
	if err != nil {
		return err
	}
	if newName != "" {
		contact.Name = newName
	}
	if newEmail != "" {
		contact.Email = newEmail
	}
	return js.save()
}

func (js *JSONStore) Delete(id int) error {
	for i, c := range js.contacts {
		if c.ID == id {
			js.contacts = append(js.contacts[:i], js.contacts[i+1:]...)
			return js.save()
		}
	}
	return ErrContactNotFound(id)
}

func (js *JSONStore) save() error {
	data, err := json.MarshalIndent(js.contacts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(js.filepath, data, 0644)
}

func (js *JSONStore) load() error {
	data, err := os.ReadFile(js.filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &js.contacts)
}
