package main

import (
	"projetContact/internal/app"
	"projetContact/internal/storage"
)

func main() {

	var store storage.Storer = storage.NewMemoryStore()
	app.Run(store)
}
