package cmd

import (
	"fmt"
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage"
	"github.com/spf13/cobra"
)

var (
	name string
	email string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajoute un nouvel utilisateur au contact.",
	Long: `Ajoute un nouvel utilisateur au contact en fournissant son nom et son email.`,

	Run: func(cmd *cobra.Command, args []string) {

		if name == "" || email == "" {
			fmt.Println("Erreur : --name et --email sont requis.")
			return
		}

		newContact := &storage.Contact{
			Name:  name,
			Email: email,
		}

		if err := store.AddContact(newContact); err != nil {
			fmt.Printf("Erreur lors de l'ajout du contact : %v\n", err)
			return
		}

		fmt.Printf("Contact ajouté : [%d] %s <%s>\n", newContact.ID, newContact.Name, newContact.Email)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&name, "name", "n", "", "Nom du contact")
	addCmd.Flags().StringVarP(&email, "email", "e", "", "Email du contact")

	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
}
