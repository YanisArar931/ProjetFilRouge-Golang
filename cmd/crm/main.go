package main

import (
	"log"

	"github.com/YanisArar931/ProjetFilRouge-Golang/cmd"
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage"
)

func main() {

	// On crée une instance du storage JSON (persistance)
	store, err := storage.NewJSONStorage("contacts.json")
	if err != nil {
		log.Fatalf("Impossible d'initialiser le storage: %v", err)
	}

	// On donne le storage à Cobra
	cmd.SetStorage(store)

	// On lance l'application CLI
	cmd.Execute()
}
