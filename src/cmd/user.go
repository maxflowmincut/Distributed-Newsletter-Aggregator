package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User-related operations",
	Long: `Perform operations related to users, such as creating a new user or listing existing ones.

Example usage:
- To create a new user: 'src user create'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'src user [subcommand]' to perform user operations.")
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(createCmd)
}