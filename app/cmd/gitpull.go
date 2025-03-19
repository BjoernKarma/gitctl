package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bjoernkarma/gitctl/config"
	"github.com/bjoernkarma/gitctl/gitrepo"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Execute git pull on multiple git repositories.",
	Run: func(cmd *cobra.Command, args []string) {
		gitrepo.RunGitCommand(gitrepo.GitPull, config.GetBaseDirs())
	},
}
