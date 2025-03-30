package cmd

import (
	"fmt"
	"log"
	"try-graphql/config"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "boilerplate",
	Short: "Welcome to boilerplate.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to boilerplate.")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(migrateCreateCmd)
	rootCmd.AddCommand(migrateDownCmd)
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateStatusCmd)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	config.Load()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
