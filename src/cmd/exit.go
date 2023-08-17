package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// exitCmd represents the exit command
var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "Terminate the CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Exiting the CLI...")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(exitCmd)
}
