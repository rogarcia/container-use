package main

import (
	"os"

	"github.com/dagger/container-use/repository"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:               "log <env>",
	Short:             "Show the log for an environment",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: suggestEnvironments,
	RunE: func(app *cobra.Command, args []string) error {
		ctx := app.Context()

		// Ensure we're in a git repository
		repo, err := repository.Open(ctx, ".")
		if err != nil {
			return err
		}

		patch, _ := app.Flags().GetBool("patch")

		return repo.Log(ctx, args[0], patch, os.Stdout)
	},
}

func init() {
	logCmd.Flags().BoolP("patch", "p", false, "Generate patch")
	rootCmd.AddCommand(logCmd)
}
