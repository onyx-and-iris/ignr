// Package cmd provides a command-line interface for generating .gitignore files.
package main

import (
	"context"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/charmbracelet/fang"
	"github.com/google/go-github/v72/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version string // Version of the CLI, set during build time

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "ignr",
	Short: "A command-line interface for generating .gitignore files",
	Long: `ignr is a command-line interface for generating .gitignore files.
It allows users to easily create and manage .gitignore files for various programming languages and frameworks.
You may also list available templates and generate .gitignore files based on those templates.`,
	SilenceUsage: true,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		var client *github.Client
		if !viper.IsSet("token") || viper.GetString("token") == "" {
			client = github.NewClient(nil)
		} else {
			client = github.NewClient(nil).WithAuthToken(viper.GetString("token"))
		}
		ctx := withClient(cmd.Context(), client)
		cmd.SetContext(ctx)
	},
	//RunE: func(cmd *cobra.Command, _ []string) error {
	//},
}

// init initialises the root command and its flags.
func init() {
	rootCmd.PersistentFlags().StringP("token", "t", "", "GitHub authentication token")
	rootCmd.PersistentFlags().IntP("height", "H", 10, "Height of the selection prompt")
	rootCmd.PersistentFlags().
		StringP("filter", "f", "startswith", "Type of filter to apply to the list of templates (e.g., 'startswith', 'contains')")
	rootCmd.PersistentFlags().BoolP("start-search", "s", false, "Start the prompt in search mode")

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix("IGNR")
	viper.AutomaticEnv()
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("height", rootCmd.PersistentFlags().Lookup("height"))
	viper.BindPFlag("filter", rootCmd.PersistentFlags().Lookup("filter"))
	viper.BindPFlag("start-search", rootCmd.PersistentFlags().Lookup("start-search"))
}

// main is the entry point of the application.
// It executes the root command and handles any errors.
func main() {
	if version == "" {
		info, ok := debug.ReadBuildInfo()
		if !ok {
			log.Fatal("could not read build info")
		}
		version = strings.Split(info.Main.Version, "-")[0]
	}

	if err := fang.Execute(context.Background(), rootCmd, fang.WithVersion(version)); err != nil {
		os.Exit(1)
	}
}
