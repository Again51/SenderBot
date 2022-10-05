package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "send",
	Short: "Envois de fichiers sur discord",
	Long: `Avec ce programme vous pouvez envoyer des fichiers sur discord
	
	Exemple:
	send file -n "nom du fichier"
	send file -n "nom du fichier" -c "ID du channel"
	send file all -p "chemin du dossier"`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
