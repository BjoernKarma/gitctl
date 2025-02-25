package cmd

import (
	"github.com/spf13/cobra"

	"ethical-developer/cli/gitctl/config"
	"ethical-developer/cli/gitctl/gitrepo"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Execute git status on multiple git repositories.",
	Run: func(cmd *cobra.Command, args []string) {
		gitrepo.RunGitCommand(gitrepo.GitStatus, config.GetBaseDirs())
	},
}
