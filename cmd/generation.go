package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/jeremyhager/pokego/internal/generation"
	"github.com/spf13/cobra"
)

var Count bool

// SpeciesCmd represents the Species command
var generationCmd = &cobra.Command{
	Use:   "generation",
	Short: "Get information about a pokemon generation.",
	Long: `generation is used for getting pokemon generation information via the pokeapi either
by id number or by name, or by using --count to get the number of pokemon generations
	
Examples:
pokego generation 1
pokego generation generation-i
pokego generation --count`,
	RunE: func(cmd *cobra.Command, args []string) error {
		genArgs := generation.GenerationArgs{
			ID:    args,
			Count: Count,
		}
		if len(args) == 0 && !genArgs.Count {
			err := cmd.Help()
			if err != nil {
				log.Fatal(err)
			}
			os.Exit(0)
		} else if len(args) == 0 && genArgs.Count {
			named, err := genArgs.GetCount()
			if err != nil {
				return err
			}
			fmt.Printf("%v\n", named.Count)
			return nil
		}
		gen, err := genArgs.Get()
		if err != nil {
			return err
		}
		fmt.Printf("gen main region: %v\n", gen.MainRegion.Name)
		return nil

	},
}

func init() {
	rootCmd.AddCommand(generationCmd)
	generationCmd.Flags().BoolVarP(&Count, "count", "c", false, "get current number of generations")
}
