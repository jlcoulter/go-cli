package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jlcoulter/go-cli-template/internal/config"
)

var (
	cfgFile string
	jsonOut bool
)

var rootCmd = &cobra.Command{
	Use:   "go-cli-template",
	Short: "A CLI tool template",
	Long:  "A CLI tool template — replace this with your tool's description.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.Load(cfgFile)
	},
}

var version string

func SetVersion(v string) {
	version = v
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default $HOME/.go-cli-template.yaml)")
	rootCmd.PersistentFlags().BoolVar(&jsonOut, "json", false, "output in JSON format")

	viper.BindPFlag("json", rootCmd.PersistentFlags().Lookup("json"))

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(exampleCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}