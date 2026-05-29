package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bjoernkarma/gitctl/config"
	"github.com/bjoernkarma/gitctl/gitrepo"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Show current branch for multiple git repositories.",
	RunE: func(cmd *cobra.Command, args []string) error {
		baseDirs, err := config.GetBaseDirs()
		if err != nil {
			return err
		}
		if err := gitrepo.RunGitCommand(gitrepo.GitBranch, baseDirs); err != nil {
			// Errors have already been displayed via the color package.
			return ErrSilent
		}
		return nil
	},
}
