package cmd

import (
	"github.com/spf13/cobra"

	"ethical-developer/cli/gitctl/config"
	"ethical-developer/cli/gitctl/gitrepo"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Execute git pull on multiple git repositories.",
	Run: func(cmd *cobra.Command, args []string) {
		gitrepo.RunGitCommand(gitrepo.GitPull, config.GetBaseDirs())
	},
}
