package main

import (
	"fmt"
	"projetContact/internal/app"
	"projetContact/internal/storage"
)

func main() {
	store, err := storage.NewGORMStore("contacts.db")
	if err != nil {
		fmt.Printf("Erreur GORM: %v\n", err)
		return
	}

	app.Run(store)
}
