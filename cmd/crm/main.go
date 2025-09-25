package main

import (
	"fmt"
	"projetContact/internal/app"
	"projetContact/internal/storage"
)

func main() {
	store, err := storage.NewJSONStore("contacts.json")
	if err != nil {
		fmt.Printf("Erreur initialisation JSON store: %v\n", err)
		return
	}

	app.Run(store)
}
