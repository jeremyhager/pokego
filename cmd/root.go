package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pokego",
	Short: "A simple cli tool to interact with pokeapi.",
	Long: `Used to showcase different aspects of coding concepts while using 
practical usecases. Ideologies and concepts include: TDD, OOP, DRY, etc.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
