package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra-demo",
	Short: "cobra-demo is an example of cobra usage",
	Long:  "cobra-demo shows how to create a CLI application with cobra",
	// If want to run bare cobra-demo:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	cobra.OnInitialize(initConfig)
	// Persistent flag is global for application
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-demo.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with ".cobra-demo" (without extension)
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra-demo")
	}
}

// Execute is primary entrypoint for all child commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
