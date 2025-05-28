package cmd

import (
	"fmt"

	"github.com/0xUmang/dictionaryapi/internal/api"
	"github.com/spf13/cobra"
)

var definitionCmd = &cobra.Command{
	Use:   "definition [word]",
	Short: "Fetches the definition of a word",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		word := args[0]

		result, err := api.FetchDefinition(ctx, word)
		if err != nil {
			return fmt.Errorf("Error fetching definition of the supplied word: %w", err)
		}

		fmt.Printf("Word Definition requested for: %s\n\n\n", word)
		if len(result.Pronunciations) > 0 {
			fmt.Println("Pronunciations:")
			for j, p := range result.Pronunciations {
				fmt.Printf(" %d. %s\n", j+1, p)
			}
		}
		fmt.Printf("\nWord Type: %s\n", result.WordType)
		fmt.Println("\nDefinitions:")
		for i, d := range result.Definitions {
			fmt.Printf(" %d. %s\n", i+1, d)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(definitionCmd)
}
