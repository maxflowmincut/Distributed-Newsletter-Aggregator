package cmd

import (
	"os"

	"newsletter-aggregator/src/db/sqlite"
	"github.com/spf13/cobra"
)

var DB *sqlite.SQLiteDB

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "src",
	Short: "Distributed Newsletter Aggregator CLI",
	Long: `Distributed Newsletter Aggregator CLI provides tools to manage and customize your newsletter subscriptions.

With this CLI:
- Sign up for personalized newsletters
- Manage user profiles
- Customize content preferences and delivery schedules

To get started, use 'src [subcommand]' or 'src --help' for more details.
Thank you for choosing Newsletter Aggregator!`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(db *sqlite.SQLiteDB) {
	DB = db 
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}