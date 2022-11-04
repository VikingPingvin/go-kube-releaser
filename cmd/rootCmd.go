package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "go-kube-release",
	Short: "Simple CLI tool to help releasing a versioned set of kubernetes resources",
	Run: func(cmd *cobra.Command, args []string) {
		println("test")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {

	}

	viper.AutomaticEnv()
}
