package cmd

import (
	"fmt"
	"os"
	"projetContact/internal/storage"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister tous les contacts",
	Long:  "Afficher la liste de tous les contacts enregistrés.",
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("format")

		store := GetStore()
		contacts, err := store.GetAll()
		if err != nil {
			return fmt.Errorf("erreur lors de la récupération des contacts: %v", err)
		}

		if len(contacts) == 0 {
			fmt.Println("Aucun contact trouvé.")
			return nil
		}

		switch format {
		case "table":
			displayContactsTable(contacts)
		case "simple":
			displayContactsSimple(contacts)
		default:
			displayContactsTable(contacts)
		}

		fmt.Printf("\n Total: %d contact(s)\n", len(contacts))
		return nil
	},
}

func displayContactsTable(contacts []*storage.Contact) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tNom\tEmail")
	fmt.Fprintln(w, "---\t---\t---")

	for _, contact := range contacts {
		fmt.Fprintf(w, "%d\t%s\t%s\n", contact.ID, contact.Name, contact.Email)
	}

	w.Flush()
}

func displayContactsSimple(contacts []*storage.Contact) {
	for _, contact := range contacts {
		fmt.Printf("[%d] %s <%s>\n", contact.ID, contact.Name, contact.Email)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringP("format", "f", "table", "Format d'affichage (table, simple)")
}
