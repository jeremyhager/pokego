/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jeremyhager/pokego/internal/generate"
	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string
var id string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates files using the Go template format",
	Long: `Using the input and output files, a template can be generated from ID.
Example:

pokego generate --pokemon 155 --input input.md.tmpl --output output.md`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := generate.RenderTemplate(inputFile, outputFile, id); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&inputFile, "input", "i", "", "input file to render (required)")
	generateCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output file to create (required)")
	generateCmd.Flags().StringVarP(&id, "id", "", "", "pokemon to base template and output on (default stdout)")
	generateCmd.MarkFlagRequired("id")
	generateCmd.MarkFlagRequired("input")
	generateCmd.MarkFlagsRequiredTogether("input", "id")
}
