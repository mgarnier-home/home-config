package cmd

import (
	"fmt"

	"mgarnier11/home-cli/utils"

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("up called")
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return utils.Test(), cobra.ShellCompDirectiveNoFileComp

	},
}

func init() {
	rootCmd.AddCommand(upCmd)
}
