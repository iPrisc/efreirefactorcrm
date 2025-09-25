package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Supprimer un contact",
	Long:  "Supprimer un contact du système en utilisant son ID.",
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

		force, _ := cmd.Flags().GetBool("force")

		if !force {
			fmt.Printf(" Êtes-vous sûr de vouloir supprimer le contact suivant ?\n")
			fmt.Printf("   [%d] %s <%s>\n", existingContact.ID, existingContact.Name, existingContact.Email)
			fmt.Print("   Tapez 'oui' pour confirmer: ")

			var confirmation string
			fmt.Scanln(&confirmation)

			if confirmation != "oui" && confirmation != "OUI" && confirmation != "o" && confirmation != "O" {
				fmt.Println("Suppression annulée")
				return nil
			}
		}

		err = store.Delete(id)
		if err != nil {
			return fmt.Errorf("erreur lors de la suppression: %v", err)
		}

		fmt.Printf(" Contact [%d] %s supprimé avec succès\n", id, existingContact.Name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().BoolP("force", "f", false, "Supprimer sans confirmation")
}
