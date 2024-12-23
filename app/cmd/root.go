package cmd

import (
	"ethical-developer/cli/gitctl/gitrepo"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	Verbose bool

	rootCmd = &cobra.Command{
		Use:   "gitctl",
		Short: "Run git commands on multiple git repositories",
		Long: `Run git commands on multiple git repositories. 
For example, you can run 'gitctl pull' to pull all the git 
repositories in the base directories.`,
		Version: "0.1.0",
	}

	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Execute git status on multiple git repositories.",
		Run: func(cmd *cobra.Command, args []string) {
			gitrepo.RunGitCommand(gitrepo.GitStatus, viper.GetStringSlice("base.dir"))
		},
	}

	pullCmd = &cobra.Command{
		Use:   "pull",
		Short: "Execute git pull on multiple git repositories.",
		Run: func(cmd *cobra.Command, args []string) {
			gitrepo.RunGitCommand(gitrepo.GitPull, viper.GetStringSlice("base.dir"))
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/gitctl.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(pullCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		log.Println("Using config file from flag:", cfgFile)
	} else {
		// Find home directory.
		var home string
		home = os.Getenv("HOME")
		if home == "" {
			var err error
			home, err = os.UserHomeDir()
			cobra.CheckErr(err)
		}

		// Search config in home directory with name "gitctl.yaml" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("gitctl")
		log.Println("Using config file from home dir:", home)
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Println("fatal error config file: ", err)
		return
	}
	log.Println("Using config file:", viper.ConfigFileUsed())
	log.Println("Found base dirs:", viper.GetStringSlice("base.dir"))
}
