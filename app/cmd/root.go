package cmd

import (
	"ethical-developer/cli/gitctl/config"
	"log"

	"github.com/pkg/errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	configFile  string
	quiet       bool
	verbose     bool
	debug       bool
	local       bool
	dryRun      bool
	color       bool
	concurrency string
	baseDirs    []string
)

var rootCmd = &cobra.Command{
	Use:   "gitctl",
	Short: "Run git commands on multiple git repositories",
	Long: `Run git commands on multiple git repositories. 
For example, you can run 'gitctl pull' to pull all the git 
repositories in the base directories.`,
	Version: "1.2.0",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(InitConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.config/gitctl.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress output")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug output")
	rootCmd.PersistentFlags().BoolVarP(&local, "local", "l", false, "run with working directory used as base directory")
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dryRun", "D", false, "run with dry run mode")
	rootCmd.PersistentFlags().BoolVarP(&color, "color", "c", true, "color output")
	rootCmd.PersistentFlags().StringVarP(&concurrency, "concurrency", "C", "1", "number of concurrent operations")
	rootCmd.PersistentFlags().StringSliceVar(&baseDirs, "base.dirs", []string{}, "base directories for git repositories")

	// Bind flags to Viper settings
	_ = viper.BindPFlag(config.GitCtlQuiet, rootCmd.PersistentFlags().Lookup("quiet"))
	_ = viper.BindPFlag(config.GitCtlVerbose, rootCmd.PersistentFlags().Lookup("verbose"))
	_ = viper.BindPFlag(config.GitCtlDebug, rootCmd.PersistentFlags().Lookup("debug"))
	_ = viper.BindPFlag(config.GitCtlLocal, rootCmd.PersistentFlags().Lookup("local"))
	_ = viper.BindPFlag(config.GitCtlDryRun, rootCmd.PersistentFlags().Lookup("dryRun"))
	_ = viper.BindPFlag(config.GitCtlColor, rootCmd.PersistentFlags().Lookup("color"))
	_ = viper.BindPFlag(config.GitCtlConcurrency, rootCmd.PersistentFlags().Lookup("concurrency"))
	_ = viper.BindPFlag(config.GitCtlBaseDirs, rootCmd.PersistentFlags().Lookup("base.dirs"))

	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(pullCmd)
}

func InitConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("gitctl")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(config.GitctlWorkingDir())
		viper.AddConfigPath(config.GitctlConfigDir())
	}

	// Enable reading from environment variables
	viper.AutomaticEnv()

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Println("No configuration file found, using defaults and environment variables")
		}
	} else {
		log.Printf("Using configuration file: %s", viper.ConfigFileUsed())
	}

	// Optionally, print the configuration settings for debugging
	if config.IsDebug() {
		log.Printf("Configuration settings: %v", viper.AllSettings())
	}
}
