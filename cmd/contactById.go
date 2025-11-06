package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var id int

var getByIdCmd = &cobra.Command{
	Use:   "getById",
	Short: "Liste un contact par son Id.",
	Long: `Liste un contact stocké dans le Mini-CRM en fournissant son Id.`,

	Run: func (cmd *cobra.Command, args[]string){
		if id  == 0 {
			fmt.Println("Erreur : veuillez fournir un ID avec --id")
			return
		}

		contact, err := store.GetByContactID(id)

		if err != nil {
			fmt.Printf("Erreur lors de la récupération du contact : %v\n", err)
			return
		}
		
		fmt.Printf("Contact trouvé : [%d] %s <%s>\n", contact.ID, contact.Name, contact.Email)
	},
}

func init() {
	rootCmd.AddCommand(getByIdCmd)
	getByIdCmd.Flags().IntVarP(&id, "id", "i", 0, "ID du contact")

	getByIdCmd.MarkFlagRequired("id")
}