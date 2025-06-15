// Package cmd provides a command-line interface for generating .gitignore files.
package cmd

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/google/go-github/v72/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version string // Version of the CLI, set during build time

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "ignr-cli",
	Short: "A command-line interface for generating .gitignore files",
	Long: `ignr-cli is a command-line interface for generating .gitignore files.
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
		ctx := context.WithValue(context.Background(), clientKey, client)
		cmd.SetContext(ctx)
	},
	RunE: func(cmd *cobra.Command, _ []string) error {
		if cmd.Flags().Lookup("version").Changed {
			if version == "" {
				info, ok := debug.ReadBuildInfo()
				if !ok {
					return errors.New("unable to retrieve build information")
				}
				version = strings.Split(info.Main.Version, "-")[0]
			}
			fmt.Printf("ignr-cli version: %s\n", version)
			return nil
		}

		return cmd.Help()
	},
}

func init() {
	rootCmd.PersistentFlags().StringP("token", "t", "", "GitHub authentication token")
	rootCmd.Flags().BoolP("version", "v", false, "Print the version of the CLI")

	viper.SetEnvPrefix("GH")
	viper.AutomaticEnv()
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
