package main

import (
	"fmt"
	"os"

	"github.com/clburlison/bakeit/client"
	"github.com/clburlison/bakeit/client/config"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "bakeit",
	Short: "bakeit is a bootstrapping tool for chef",
	Long: `bakeit is a multi platform bootstrapping tool used to
to install, configure, and run chef-client during bootstrap on nodes.

Complete documentation is available at https://github.com/clburlison/bakeit/.`,
	Run: func(cmd *cobra.Command, args []string) {
		client.Setup()
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&config.Verbose, "verbose", "v", true, "verbose output")
	RootCmd.Flags().BoolP("clitools", "", false, "install Xcode cli tools (not yet implemented)")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
