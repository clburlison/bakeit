package main

import (
	"github.com/clburlison/bakeit/client/version"
	"github.com/spf13/cobra"
)

var fFull bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of bakeit",
	Long:  `Print the version and build information of bakeit`,
	Run: func(cmd *cobra.Command, args []string) {
		if fFull {
			version.PrintFull()
			return
		}
		version.Print()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	versionCmd.PersistentFlags().BoolVar(&fFull, "full", false, "print full version information")
}
