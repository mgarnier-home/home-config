package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")
	},
	ValidArgs: []string{"foo", "bar"},
}

func init() {
	rootCmd.AddCommand(pullCmd)
}
