package species

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/jeremyhager/pokeapi"
	"github.com/spf13/cobra"
)

type SpeciesOptions struct {
	SpeciesID                              string
	Generation, ID, Name, TextEntries, raw bool
}

func NewSpeciesCmd() *cobra.Command {
	opts := &SpeciesOptions{}
	cmd := &cobra.Command{
		Use:   "species",
		Short: "Get information about a pokemon species.",
		Long: heredoc.Doc(`
			The species command is used for getting pokemon species information via the pokeapi either by id number or by name.
			At least 1 flag must be specified.
		`),
		Example: heredoc.Doc(`
		$ pokego species 155 --generation # get generation name by id
		$ pokego species cyndaquil --generation # get generation name by name
		`),

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.SpeciesID = args[0]
			} else if len(args) == 0 {
				err := cmd.Help()
				if err != nil {
					log.Fatal(err)
				}
				os.Exit(0)
			}
			return speciesRun(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.Generation, "generation", "g", false, "Get generation id of a species")
	cmd.Flags().BoolVarP(&opts.ID, "id", "", false, "Get ID of a species")
	cmd.Flags().BoolVarP(&opts.Name, "name", "n", false, "Get name of species")
	cmd.Flags().BoolVarP(&opts.TextEntries, "entries", "e", false, "Get a list of flavor text entries for this Pokémon species.")
	cmd.Flags().BoolVarP(&opts.raw, "raw", "", false, "Get the raw json response.")

	return cmd
}

func speciesRun(opts *SpeciesOptions) error {
	species, err := pokeapi.GetSpecies(opts.SpeciesID)
	if err != nil {
		return err
	}

	if !opts.Generation && !opts.ID && !opts.Name && !opts.TextEntries && !opts.raw {
		return fmt.Errorf("at least 1 flag must be specified")
	}

	if opts.Generation {
		rawGeneration, _ := json.Marshal(species.Generation)
		fmt.Printf("%s", rawGeneration)
	}
	if opts.ID {
		fmt.Printf("%v", species.ID)
	}
	if opts.Name {
		fmt.Printf("%v", species.Name)
	}
	if opts.TextEntries {
		rawTextEntries, _ := json.Marshal(species.FlavorTextEntries)
		fmt.Printf("%s", rawTextEntries)
	}
	if opts.raw {
		rawOutput, _ := json.Marshal(species)
		fmt.Printf("%s", rawOutput)
	}

	return nil
}
