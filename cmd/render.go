/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jeremyhager/pokego/internal/render"
	"github.com/spf13/cobra"
)

var inputFile string
var outputFile string
var id string

// generateCmd represents the generate command
var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Renders files using the Go template format",
	Long: `Using the input and output files, a template can be generated from ID.
Example:

pokego generate --pokemon 155 --input pokemon.md.tmpl --output 155.md`,
	RunE: func(cmd *cobra.Command, args []string) error {
		renderArgs := render.RenderArgs{
			ID:         id,
			InputFile:  inputFile,
			OutputFile: outputFile,
		}
		if err := renderArgs.RenderTemplate(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(renderCmd)

	renderCmd.Flags().StringVarP(&inputFile, "input", "i", "", "input file to render (required)")
	renderCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output file to create (default stdout)")
	renderCmd.Flags().StringVarP(&id, "id", "", "", "pokemon to base template and output on (required)")
	renderCmd.MarkFlagRequired("id")
	renderCmd.MarkFlagRequired("input")
	renderCmd.MarkFlagsRequiredTogether("input", "id")
}
