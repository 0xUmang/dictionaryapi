package cmd

import (
    "context"
    "github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
    Use:   "DictionaryApi",
    Short: "get word definitions",	
	Long: `A CLI tool to fetch word definitions from merriam-webster dictionary API `,
}

func Execute(ctx context.Context) error {
    return rootCmd.ExecuteContext(ctx)
}