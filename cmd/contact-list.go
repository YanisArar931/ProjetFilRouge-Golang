package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "Lister tous les contacts.",
	Long: `Lister tous les contacts stockés dans le Mini-CRM.`,

	Run: func (cmd *cobra.Command, args[]string){
		contacts, err := store.GetAllContacts()

		if err != nil {
			fmt.Printf("Erreur lors de la récupération des contacts : %v\n", err)
			return
		}

		if len(contacts) == 0 {
			fmt.Println("Aucun contact trouvé.")
			return
		}
		
		fmt.Println("Liste des contacts :")
		for _, contact := range contacts {
			fmt.Printf("[%d] %s <%s>\n", contact.ID, contact.Name, contact.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(getAllCmd)
}