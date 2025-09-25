package cmd

import (
	"fmt"
	"os"
	"projetContact/internal/storage"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	store   storage.Storer
)

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "Un gestionnaire de contacts simple et efficace",
	Long:  `Mini-CRM CLI - Gestionnaire de contacts en ligne de commande`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "fichier de configuration (défaut: ./config.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("Utilisation du fichier de configuration: %s\n", viper.ConfigFileUsed())
	} else {
		fmt.Printf("Impossible de lire le fichier de configuration: %v\n", err)
		fmt.Println("Utilisation des valeurs par défaut...")
	}

	initStore()
}

func initStore() {
	storageType := viper.GetString("storage.type")

	switch storageType {
	case "memory":
		store = storage.NewMemoryStore()
		fmt.Println("Utilisation du stockage en mémoire")
	case "json":
		jsonFile := viper.GetString("storage.json.file")
		if jsonFile == "" {
			jsonFile = "contacts.json"
		}
		var err error
		store, err = storage.NewJSONStore(jsonFile)
		if err != nil {
			fmt.Printf("Erreur lors de l'initialisation du stockage JSON: %v\n", err)
			fmt.Println("Utilisation du stockage en mémoire comme fallback")
			store = storage.NewMemoryStore()
		} else {
			fmt.Printf("Utilisation du stockage JSON: %s\n", jsonFile)
		}
	case "gorm":
		dbFile := viper.GetString("storage.gorm.database")
		if dbFile == "" {
			dbFile = "contacts.db"
		}
		var err error
		store, err = storage.NewGORMStore(dbFile)
		if err != nil {
			fmt.Printf("Erreur lors de l'initialisation du stockage GORM: %v\n", err)
			fmt.Println("Utilisation du stockage en mémoire comme fallback")
			store = storage.NewMemoryStore()
		} else {
			fmt.Printf("Utilisation du stockage GORM/SQLite: %s\n", dbFile)
		}
	default:
		fmt.Printf("Type de stockage inconnu '%s', utilisation du stockage en mémoire\n", storageType)
		store = storage.NewMemoryStore()
	}
}

func GetStore() storage.Storer {
	return store
}
