package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [ID]",
	Short: "Mettre à jour un contact existant",
	Long:  "Mettre à jour le nom et/ou l'email d'un contact existant.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Conversion de l'ID
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("ID invalide '%s': doit être un nombre", args[0])
		}

		store := GetStore()

		existingContact, err := store.GetByID(id)
		if err != nil {
			return fmt.Errorf("contact avec l'ID %d non trouvé", id)
		}

		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		if name == "" {
			name = existingContact.Name
		}
		if email == "" {
			email = existingContact.Email
		}

		if name == existingContact.Name && email == existingContact.Email {
			fmt.Printf(" Aucune modification détectée pour le contact [%d] %s\n", id, existingContact.Name)
			return nil
		}

		err = store.Update(id, name, email)
		if err != nil {
			return fmt.Errorf("erreur lors de la mise à jour: %v", err)
		}

		fmt.Printf("Contact [%d] mis à jour avec succès\n", id)
		fmt.Printf("   Nom: %s\n", name)
		fmt.Printf("   Email: %s\n", email)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("name", "n", "", "Nouveau nom du contact")
	updateCmd.Flags().StringP("email", "e", "", "Nouvel email du contact")
}
