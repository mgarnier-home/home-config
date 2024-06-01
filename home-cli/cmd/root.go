package cmd

import (
	"fmt"
	"mgarnier11/home-cli/command"
	"mgarnier11/home-cli/utils"
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "home-cli",
}
var stackList = utils.GetStacks()
var hostList = utils.GetHosts()
var actionList = utils.GetActions()

func getCommandArgs(commandPath string) (stack string, host string, action string) {
	parts := strings.Split(commandPath, " ")

	if len(parts) == 1 {
		return "", "", ""
	} else if len(parts) == 2 {
		return "all", "all", parts[1]
	} else if len(parts) == 3 {

		if slices.Contains(stackList, parts[1]) {
			return parts[1], "all", parts[2]
		} else {
			return "all", parts[1], parts[2]
		}

	}

	return parts[1], parts[2], parts[3]

}

func createActionsCommands() []*cobra.Command {
	actionsCmd := []*cobra.Command{}

	for _, action := range actionList {
		actionCmd := &cobra.Command{
			Use:       action,
			ValidArgs: []string{"parallel"},
			Args:      cobra.MatchAll(cobra.RangeArgs(0, 1), cobra.OnlyValidArgs),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Running action", cmd.CommandPath(), "with args ", args)

				stack, host, action := getCommandArgs(cmd.CommandPath())

				command.ExecCommand(stack, host, action, args)
			},
		}

		actionsCmd = append(actionsCmd, actionCmd)
	}

	return actionsCmd
}

func createHostsCommands(stack string) []*cobra.Command {
	hosts := hostList

	if stack != "" {
		hosts = utils.GetHostsByStack(stack)
	}

	hostsCmd := []*cobra.Command{}

	for _, host := range hosts {
		hostCmd := &cobra.Command{
			Use: host,
		}

		hostCmd.AddCommand(createActionsCommands()...)
		hostsCmd = append(hostsCmd, hostCmd)
	}

	return hostsCmd
}

func createStacksCommands() []*cobra.Command {
	stacksCmd := []*cobra.Command{}

	for _, stack := range stackList {
		stackCmd := &cobra.Command{
			Use: stack,
		}

		stackCmd.AddCommand(createHostsCommands(stack)...)
		stackCmd.AddCommand(createActionsCommands()...)
		stacksCmd = append(stacksCmd, stackCmd)
	}

	return stacksCmd
}

func Execute() {
	rootCmd.AddCommand(createStacksCommands()...)
	rootCmd.AddCommand(createHostsCommands("")...)
	rootCmd.AddCommand(createActionsCommands()...)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.CompletionOptions.DisableDescriptions = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}
