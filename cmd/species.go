package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jeremyhager/pokeapi"
	"github.com/spf13/cobra"
)

// SpeciesCmd represents the Species command
var SpeciesCmd = &cobra.Command{
	Use:   "species",
	Short: "Get information about a pokemon species.",
	Long: `species command is used for getting pokemon species information via the pokeapi either
by id number or by name.
	
Examples:
pokego species 1
pokego species bulbasaur`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		}

		species, err := pokeapi.GetSpecies(args[0])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("pokemon species description:\n\n%v\n", species.FlavorTextEntries[0].FlavorText)
	},
}

func init() {
	rootCmd.AddCommand(SpeciesCmd)
}
