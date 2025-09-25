package cmd

import (
	"fmt"
	"projetContact/internal/storage"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un nouveau contact",
	Long:  "Ajouter un nouveau contact au système.",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		if name == "" {
			return fmt.Errorf("le nom est requis (utilisez --name ou -n)")
		}
		if email == "" {
			return fmt.Errorf("l'email est requis (utilisez --email ou -e)")
		}

		contact := &storage.Contact{
			Name:  name,
			Email: email,
		}

		store := GetStore()
		err := store.Add(contact)
		if err != nil {
			return fmt.Errorf("erreur lors de l'ajout du contact: %v", err)
		}

		fmt.Printf("Contact '%s' ajouté avec succès avec l'ID %d\n", contact.Name, contact.ID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("name", "n", "", "Nom du contact (requis)")
	addCmd.Flags().StringP("email", "e", "", "Email du contact (requis)")

	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
}
