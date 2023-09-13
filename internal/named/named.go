package named

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/jeremyhager/pokeapi"
	"github.com/spf13/cobra"
)

type NamedOptions struct {
	Endpoint                            string
	Count, Next, Previous, Results, raw bool
}

func NewNamedCmd() *cobra.Command {
	opts := &NamedOptions{}
	cmd := &cobra.Command{
		Use:   "named",
		Short: "Get any API endpoint without a resource ID or name.",
		Long: heredoc.Doc(`
			The named command prints all information about the named pokeAPI endpoint.
			At least 1 flag must be specified.
		`),
		Example: heredoc.Doc(`
			$ pokego named pokemon # get all info about the pokemon endpoint
			$ pokego named pokemon --count # get count from result
		`),

		PreRunE: cobra.ArbitraryArgs,

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				opts.Endpoint = args[0]
			} else if len(args) == 0 {
				err := cmd.Help()
				if err != nil {
					log.Fatal(err)
				}
				os.Exit(0)
			}
			return namedRun(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.Count, "count", "c", false, "Get the count of a resource")
	cmd.Flags().BoolVarP(&opts.Next, "next", "n", false, "Get the next page of a resource")
	cmd.Flags().BoolVarP(&opts.Previous, "previous", "p", false, "Get the previous of a resource")
	cmd.Flags().BoolVarP(&opts.Results, "results", "r", false, "Get the results of a resource")
	cmd.Flags().BoolVarP(&opts.raw, "raw", "", false, "Get the raw json response of a resource")

	return cmd
}

func namedRun(opts *NamedOptions) error {
	named, err := pokeapi.GetNamedEndpoint(opts.Endpoint)
	if err != nil {
		return err
	}

	if !opts.Count && !opts.Next && !opts.Previous && !opts.Results && !opts.raw {
		return fmt.Errorf("at least 1 flag must be specified")
	}

	if opts.Count {
		fmt.Printf("%v\n", named.Count)
	}
	if opts.Next {
		fmt.Printf("%v\n", named.Next)
	}
	if opts.Previous {
		fmt.Printf("%v\n", named.Previous)
	}
	if opts.Results {
		rawResults, _ := json.Marshal(named.Results)
		fmt.Printf("%s", rawResults)
	}

	if opts.raw {
		rawOutput, _ := json.Marshal(named)
		fmt.Printf("%s", rawOutput)
	}

	return nil
}
