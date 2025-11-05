package main

import (
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/app"
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	app.Run(store)
}
