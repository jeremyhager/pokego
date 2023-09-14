package cmd

import (
	"os"

	"github.com/jeremyhager/pokego/internal/named"
	"github.com/jeremyhager/pokego/internal/pokemon"
	"github.com/jeremyhager/pokego/internal/species"
	"github.com/spf13/cobra"
)

var (
	speciesCmd = species.NewSpeciesCmd()
	pokemonCmd = pokemon.NewPokemonCmd()
	namedCmd   = named.NewNamedCmd()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pokego",
	Short: "A simple cli tool to interact with pokeapi.",
	Long: `Used to showcase different aspects of coding concepts while using 
practical usecases. Ideologies and concepts include: TDD, OOP, DRY, etc.`,
	SilenceUsage: true,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(namedCmd)
	rootCmd.AddCommand(speciesCmd)
	rootCmd.AddCommand(pokemonCmd)

	// rootCmd.PersistentFlags().String("language", "en", "change the output language")
	// rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "", false, "debug output")
}
