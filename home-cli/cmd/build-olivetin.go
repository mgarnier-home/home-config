package cmd

import (
	"fmt"
	"mgarnier11/home-cli/command"
	"mgarnier11/home-cli/utils"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

var buildOlivetinCmd = &cobra.Command{
	Use:   "build-olivetin",
	Short: "Build olivetin config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building olivetin config")

		commandsPaths := getSubCommandsPaths(command.GetCobraCommands())

		globalStackControlCommands := []string{}
		stacksControlCommands := []string{}
		hostControlCommands := []string{}
		actionCommands := []string{}

		for _, commandPath := range commandsPaths {
			commandParts := strings.Split(commandPath, " ")

			switch len(commandParts) {
			case 1:
				actionCommands = append(actionCommands, commandPath)
			case 2:
				if slices.Contains(utils.StackList, commandParts[0]) {
					stacksControlCommands = append(stacksControlCommands, commandPath)
				} else if slices.Contains(utils.HostList, commandParts[0]) {
					hostControlCommands = append(hostControlCommands, commandPath)
				}
			case 3:
				globalStackControlCommands = append(globalStackControlCommands, commandPath)
			}
		}

		fmt.Println("===========Global stack control commands===========")
		fmt.Println(strings.Join(globalStackControlCommands, "\n"))

		fmt.Println("===========Stacks control commands===========")
		fmt.Println(strings.Join(stacksControlCommands, "\n"))

		fmt.Println("===========Host control commands===========")
		fmt.Println(strings.Join(hostControlCommands, "\n"))

		fmt.Println("===========Action commands===========")
		fmt.Println(strings.Join(actionCommands, "\n"))

	},
}

func getSubCommandsPaths(commands []*cobra.Command) []string {
	paths := []string{}

	for _, command := range commands {
		if slices.Contains(utils.ActionList, command.Use) {
			paths = append(paths, command.CommandPath())
		}

		paths = append(paths, getSubCommandsPaths(command.Commands())...)
	}

	return paths
}

func init() {
	buildOlivetinCmd.Hidden = true
	rootCmd.AddCommand(buildOlivetinCmd)
}
