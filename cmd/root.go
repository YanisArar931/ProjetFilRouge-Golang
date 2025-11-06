package cmd

import (
	"fmt"
	"os"

	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage"
	"github.com/spf13/cobra"
)
var store storage.Storer

var rootCmd = &cobra.Command{
	Use: "mini-crm",
	Short: "Une application Mini-CRM en ligne de commande",
	Long: ``,
}

func Execute () {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func init() {
	store = storage.NewMemoryStore()
}