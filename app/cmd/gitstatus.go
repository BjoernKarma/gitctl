package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bjoernkarma/gitctl/config"
	"github.com/bjoernkarma/gitctl/gitrepo"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Execute git status on multiple git repositories.",
	RunE: func(cmd *cobra.Command, args []string) error {
		baseDirs, err := config.GetBaseDirs()
		if err != nil {
			return err
		}
		if err := gitrepo.RunGitCommand(gitrepo.GitStatus, baseDirs); err != nil {
			// Errors have already been displayed via the color package.
			return ErrSilent
		}
		return nil
	},
}
