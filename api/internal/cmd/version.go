package cmd

import (
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  "Print program version.",
	Run:   versionExec,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func versionExec(cmd *cobra.Command, args []string) {
	println("v0")
}
