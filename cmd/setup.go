package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func setup() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	todo := filepath.Join(home, ".todo")
	err = os.MkdirAll(todo, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}
}

// setupCmd represents the setup command
var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "setup the .todo directory",
	Long:  "setup the .todo directory",
	Run: func(cmd *cobra.Command, args []string) {
		setup()
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
