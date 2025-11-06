package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Permet de mettre à jour un contact par son Id.",
	Long: `Permet de mettre à jour un contact stocké dans le Mini-CRM en fournissant son Id.`,

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
		
		fmt.Printf("Contact trouvé pour mise à jour : [%d] %s <%s>\n", contact.ID, contact.Name, contact.Email)

		fmt.Print("Entrez le nouveau nom : ")
		var newName string
		fmt.Scan(&newName)

		fmt.Print("Entrez le nouvel email : ")
		var newEmail string
		fmt.Scan(&newEmail)

		contact.Name = newName
		contact.Email = newEmail

		err = store.UpdateContact(contact.ID, contact.Name, contact.Email)
		if err != nil {
			fmt.Printf("Erreur lors de la mise à jour du contact : %v\n", err)
			return
		}

		fmt.Printf("Contact mis à jour : [%d] %s <%s>\n", contact.ID, contact.Name, contact.Email)
	},

	
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().IntVarP(&id, "id", "i", 0, "ID du contact")

	updateCmd.MarkFlagRequired("id")
}