package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bjoernkarma/gitctl/config"
	"github.com/bjoernkarma/gitctl/gitrepo"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Execute git fetch on multiple git repositories.",
	RunE: func(cmd *cobra.Command, args []string) error {
		baseDirs, err := config.GetBaseDirs()
		if err != nil {
			return err
		}
		if err := gitrepo.RunGitCommand(gitrepo.GitFetch, baseDirs); err != nil {
			// Errors have already been displayed via the color package.
			return ErrSilent
		}
		return nil
	},
}
