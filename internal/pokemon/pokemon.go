package pokemon

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/jeremyhager/pokeapi"
	"github.com/spf13/cobra"
)

type PokemonOptions struct {
	PokemonID                                   string
	id, name, raw, offcialArtwork, frontDefault bool
}

func NewPokemonCmd() *cobra.Command {
	opts := &PokemonOptions{}
	cmd := &cobra.Command{
		Use:   "pokemon",
		Short: "Get any API endpoint without a resource ID or name.",
		Long: heredoc.Doc(`
		The pokemon command is used for getting pokemon via the pokeapi eitherby id number or by name.
		
		At least 1 flag must be specified.
		`),
		Example: heredoc.Doc(`
			$ pokego pokemon 155 --name # get name for pokemon with id 155
			$ pokego species cyndaquil --official-artwork # get official artwork link for cyndaquil
		`),

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.PokemonID = args[0]
			} else if len(args) == 0 {
				err := cmd.Help()
				if err != nil {
					log.Fatal(err)
				}
				os.Exit(0)
			}
			return pokemonRun(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.id, "id", "", false, "Get ID of a species")
	cmd.Flags().BoolVarP(&opts.name, "name", "n", false, "Get name of species")
	cmd.Flags().BoolVarP(&opts.offcialArtwork, "official-artwork", "", false, "Get the link to official, non-shiny artwork")
	cmd.Flags().BoolVarP(&opts.frontDefault, "front-default", "", false, "Get the link to default, front, non-shiny sprite")
	cmd.Flags().BoolVarP(&opts.raw, "raw", "", false, "Get the raw json response.")

	return cmd
}

func pokemonRun(opts *PokemonOptions) error {
	pokemon, err := pokeapi.GetPokemon(opts.PokemonID)
	if err != nil {
		return err
	}

	if !opts.id && !opts.name && !opts.raw && !opts.offcialArtwork && !opts.frontDefault {
		return fmt.Errorf("at least 1 flag must be specified")
	}

	if opts.id {
		fmt.Printf("%v", pokemon.ID)
	}
	if opts.name {
		fmt.Printf("%v\n", pokemon.Name)
	}
	if opts.raw {
		rawOutput, _ := json.Marshal(pokemon)
		fmt.Printf("%s", rawOutput)
	}
	if opts.offcialArtwork {
		fmt.Printf("%v\n", pokemon.Sprites.Other.OfficialArtwork.FrontDefault)
	}
	if opts.frontDefault {
		fmt.Printf("%v\n", pokemon.Sprites.FrontDefault)
	}

	return nil
}
