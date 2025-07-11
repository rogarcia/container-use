package main

import (
	"fmt"

	"github.com/dagger/container-use/repository"
	"github.com/spf13/cobra"
)

var checkoutCmd = &cobra.Command{
	Use:               "checkout <env>",
	Short:             "Check out an environment in git",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: suggestEnvironments,
	RunE: func(app *cobra.Command, args []string) error {
		ctx := app.Context()
		envID := args[0]

		// Ensure we're in a git repository
		repo, err := repository.Open(ctx, ".")
		if err != nil {
			return err
		}

		branchName, err := app.Flags().GetString("branch")
		if err != nil {
			return err
		}

		branch, err := repo.Checkout(ctx, envID, branchName)
		if err != nil {
			return err
		}

		fmt.Printf("Switched to branch '%s'\n", branch)
		return nil
	},
}

func init() {
	checkoutCmd.Flags().StringP("branch", "b", "", "Local branch name to use")
	rootCmd.AddCommand(checkoutCmd)
}
