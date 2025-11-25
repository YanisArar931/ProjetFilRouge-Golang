package cmd

import (
	"fmt"

	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage"
	"github.com/YanisArar931/ProjetFilRouge-Golang/internal/storage/gormstore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var store storage.Storer

var rootCmd = &cobra.Command{
	Use:   "mini-crm",
	Short: "Mini CRM CLI en Go",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Chargement de la config YAML
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("config.yaml chargé :", viper.ConfigFileUsed())
		} else {
			fmt.Println("config.yaml introuvable, fallback sur JSON")
		}

		storeType := viper.GetString("store")
		jsonPath := viper.GetString("json_path")
		dbPath := viper.GetString("db_path")

		switch storeType {
		case "GORM":
			db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
			if err != nil {
				return err
			}
			store = gormstore.NewGORMStore(db)
			fmt.Println("Storage GORM activé")
		case "DUAL":
			jsonStore, err := storage.NewJSONStorage(jsonPath)
			if err != nil {
				return err
			}
			db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
			if err != nil {
				return err
			}
			gormStore := gormstore.NewGORMStore(db)
			store = storage.NewDualStore(jsonStore, gormStore)
			fmt.Println("Storage JSON + GORM (DUAL) activé")
		case "MEMORY":
			store = storage.NewMemoryStorage()
			fmt.Println("Storage en mémoire activé")
		default:
			jsonStore, err := storage.NewJSONStorage(jsonPath)
			if err != nil {
				return err
			}
			store = jsonStore
			fmt.Println("Storage JSON par défaut activé")
		}
		return nil
	},
}

func SetStorage(s storage.Storer) {
	store = s
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().String("store", "", "Type de storage: JSON, GORM, DUAL, MEMORY")
	rootCmd.PersistentFlags().String("json_path", "", "Chemin du fichier JSON")
	rootCmd.PersistentFlags().String("db_path", "", "Chemin du fichier SQLite DB")

	viper.BindPFlag("store", rootCmd.PersistentFlags().Lookup("store"))
	viper.BindPFlag("json_path", rootCmd.PersistentFlags().Lookup("json_path"))
	viper.BindPFlag("db_path", rootCmd.PersistentFlags().Lookup("db_path"))

	viper.SetDefault("store", "JSON")
	viper.SetDefault("json_path", "contacts.json")
	viper.SetDefault("db_path", "contacts.db")
}
