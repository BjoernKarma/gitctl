package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bjoernkarma/gitctl/config"
	"github.com/bjoernkarma/gitctl/gitrepo"
)

var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "Execute git stash on multiple git repositories.",
	RunE: func(cmd *cobra.Command, args []string) error {
		baseDirs, err := config.GetBaseDirs()
		if err != nil {
			return err
		}
		if err := gitrepo.RunGitCommand(gitrepo.GitStash, baseDirs); err != nil {
			// Errors have already been displayed via the color package.
			return ErrSilent
		}
		return nil
	},
}
