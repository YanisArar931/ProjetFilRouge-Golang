package main

import (
	"log"

	"github.com/YanisArar931/ProjetFilRouge-Golang/cmd"
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage"
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage/gormstore"
)

func main() {
	jsonStore, err := storage.NewJSONStorage("contacts.json")
	if err != nil {
		log.Fatal("Erreur JSONStorage:", err)
	}

	gormStore, err := gormstore.NewStore("contacts.db")
	if err != nil {
		log.Fatal("Erreur GORMStorage:", err)
	}

	store := storage.NewDualStore(jsonStore, gormStore)

	cmd.SetStorage(store)

	cmd.Execute()
}
