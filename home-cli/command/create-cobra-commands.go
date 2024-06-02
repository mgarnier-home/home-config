package command

import (
	"fmt"
	"mgarnier11/home-cli/utils"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

func getCommandArgs(commandPath string) (stack string, host string, action string) {
	parts := strings.Split(commandPath, " ")

	if len(parts) == 1 {
		return "", "", ""
	} else if len(parts) == 2 {
		return "all", "all", parts[1]
	} else if len(parts) == 3 {

		if slices.Contains(utils.StackList, parts[1]) {
			return parts[1], "all", parts[2]
		} else {
			return "all", parts[1], parts[2]
		}

	}

	return parts[1], parts[2], parts[3]

}

func createActionsCommands() []*cobra.Command {
	actionsCmd := []*cobra.Command{}

	for _, action := range utils.ActionList {
		actionCmd := &cobra.Command{
			Use:       action,
			ValidArgs: []string{"parallel"},
			Args:      cobra.MatchAll(cobra.RangeArgs(0, 1), cobra.OnlyValidArgs),
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Running action", cmd.CommandPath(), "with args ", args)

				stack, host, action := getCommandArgs(cmd.CommandPath())

				ExecCommand(stack, host, action, args)
			},
		}

		actionsCmd = append(actionsCmd, actionCmd)
	}

	return actionsCmd
}

func createHostsCommands(stack string) []*cobra.Command {
	hosts := utils.HostList

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

	for _, stack := range utils.StackList {
		stackCmd := &cobra.Command{
			Use: stack,
		}

		stackCmd.AddCommand(createHostsCommands(stack)...)
		stackCmd.AddCommand(createActionsCommands()...)
		stacksCmd = append(stacksCmd, stackCmd)
	}

	return stacksCmd
}

func GetCobraCommands() []*cobra.Command {
	stacksCmd := []*cobra.Command{}

	stacksCmd = append(stacksCmd, createStacksCommands()...)
	stacksCmd = append(stacksCmd, createHostsCommands("")...)
	stacksCmd = append(stacksCmd, createActionsCommands()...)

	return stacksCmd
}
