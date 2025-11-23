package cmd

import (
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mini-crm",
	Short: "Une application Mini-CRM en ligne de commande",
}

var store storage.Storer

// Appelé par main.go
func SetStorage(s storage.Storer) {
	store = s
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
