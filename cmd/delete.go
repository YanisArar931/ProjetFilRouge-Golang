package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprimer utilisateur des contacts.",
	Long: `Supprimer un utilisateur des contacts en fournissant son Id.`,

	Run: func(cmd *cobra.Command, args []string) {
		if id  == 0 {
			fmt.Println("Erreur : veuillez fournir un ID avec --id")
			return
		}

		err := store.DeleteContact(id)
		if err != nil {
			fmt.Printf("Erreur lors de la suppression du contact : %v\n", err)
			return
		}

		fmt.Printf("Contact avec l'ID %d supprimé avec succès.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "ID du contact à supprimer")

	deleteCmd.MarkFlagRequired("id")
}
