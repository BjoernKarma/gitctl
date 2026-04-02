package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bjoernkarma/gitctl/config"
	"github.com/bjoernkarma/gitctl/gitrepo"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Execute git pull on multiple git repositories.",
	RunE: func(cmd *cobra.Command, args []string) error {
		baseDirs, err := config.GetBaseDirs()
		if err != nil {
			return err
		}
		return gitrepo.RunGitCommand(gitrepo.GitPull, baseDirs)
	},
}
