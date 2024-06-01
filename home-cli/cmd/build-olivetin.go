package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var buildOlivetinCmd = &cobra.Command{
	Use:   "build-olivetin",
	Short: "Build olivetin folder",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building olivetin folder")
	},
}

func init() {
	buildOlivetinCmd.Hidden = true
	rootCmd.AddCommand(buildOlivetinCmd)
}
