package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")
	},
	ValidArgs: []string{"foo", "bar"},
}

func init() {
	rootCmd.AddCommand(downCmd)
}
