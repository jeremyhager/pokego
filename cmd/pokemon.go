package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jeremyhager/pokeapi/pokemon"
	"github.com/spf13/cobra"
)

// pokemonCmd represents the pokemon command
var pokemonCmd = &cobra.Command{
	Use:   "pokemon",
	Short: "A command to get pokemon.",
	Long: `pokemon command is used for getting pokemon via the pokeapi either
by id number or by name.

Examples:
pokego pokemon 1
pokego pokemon bulbasaur
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		poke, err := pokemon.Get(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("pokemon info:\n%+v\n", poke.Species)
	},
}

func init() {
	rootCmd.AddCommand(pokemonCmd)
}
