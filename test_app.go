package main

import (
	"fmt"
	"projetContact/internal/app"
	"projetContact/internal/storage"
)

func main() {
	// Test d'abord avec le stockage en mémoire pour vérifier que l'app fonctionne
	fmt.Println("=== Test avec MemoryStore ===")
	memStore := storage.NewMemoryStore()
	
	// Ajout d'un contact de test
	testContact := &storage.Contact{
		Name:  "Test User",
		Email: "test@example.com",
	}
	
	err := memStore.Add(testContact)
	if err != nil {
		fmt.Printf("Erreur ajout: %v\n", err)
		return
	}
	
	// Récupération des contacts
	contacts, err := memStore.GetAll()
	if err != nil {
		fmt.Printf("Erreur récupération: %v\n", err)
		return
	}
	
	fmt.Printf("Nombre de contacts en mémoire: %d\n", len(contacts))
	if len(contacts) > 0 {
		fmt.Printf("Premier contact: ID=%d, Nom=%s, Email=%s\n", 
			contacts[0].ID, contacts[0].Name, contacts[0].Email)
	}
	
	fmt.Println("\n=== Test avec JSONStore ===")
	jsonStore, err := storage.NewJSONStore("test_contacts.json")
	if err != nil {
		fmt.Printf("Erreur JSONStore: %v\n", err)
		return
	}
	
	// Ajout d'un contact de test dans JSON
	testContact2 := &storage.Contact{
		Name:  "JSON User",
		Email: "json@example.com",
	}
	
	err = jsonStore.Add(testContact2)
	if err != nil {
		fmt.Printf("Erreur ajout JSON: %v\n", err)
		return
	}
	
	contacts, err = jsonStore.GetAll()
	if err != nil {
		fmt.Printf("Erreur récupération JSON: %v\n", err)
		return
	}
	
	fmt.Printf("Nombre de contacts en JSON: %d\n", len(contacts))
	
	fmt.Println("\n=== Test avec GORMStore ===")
	gormStore, err := storage.NewGORMStore("test_contacts.db")
	if err != nil {
		fmt.Printf("Erreur GORM: %v\n", err)
		fmt.Println("GORM ne fonctionne pas, mais les autres storages fonctionnent!")
		return
	}
	
	// Ajout d'un contact de test dans GORM
	testContact3 := &storage.Contact{
		Name:  "GORM User",
		Email: "gorm@example.com",
	}
	
	err = gormStore.Add(testContact3)
	if err != nil {
		fmt.Printf("Erreur ajout GORM: %v\n", err)
		return
	}
	
	contacts, err = gormStore.GetAll()
	if err != nil {
		fmt.Printf("Erreur récupération GORM: %v\n", err)
		return
	}
	
	fmt.Printf("Nombre de contacts en GORM: %d\n", len(contacts))
	fmt.Println("GORM fonctionne parfaitement!")
	
	fmt.Println("\n=== Lancement de l'application interactive ===")
	// Lancement de l'app avec GORM
	app.Run(gormStore)
}