package cmd

import (
	"fmt"
	"mgarnier11/home-cli/compose"
	"mgarnier11/home-cli/utils"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

type action struct {
	title string
	icon  string
	shell string
}

func createActionCommands(stack string, host string) []*action {
	actions := []*action{}

	for _, actionStr := range utils.ActionList {

		shell := "home-cli"

		if stack != "all" {
			shell += " " + stack
		}
		if host != "all" {
			shell += " " + host
		}
		shell += " " + actionStr

		actions = append(actions, &action{
			title: stack + " " + host + " " + actionStr,
			icon:  "box",
			shell: shell,
		})
	}

	return actions
}

func printAction(action *action) {
	fmt.Println(action.title + " => " + action.shell)
}

func printActions(actions []*action) {
	for _, action := range actions {
		printAction(action)
	}
}

var buildOlivetinCmd = &cobra.Command{
	Use:   "build-olivetin",
	Short: "Build olivetin config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Building olivetin config")

		config := compose.GetConfig()

		actions := []*action{}

		for stack, hosts := range config.Stacks {
			if len(hosts) > 1 {
				actions = append(actions, createActionCommands(stack, "all")...)
			}

			for _, host := range hosts {
				actions = append(actions, createActionCommands(stack, host)...)
			}
		}

		for host, stacks := range config.Hosts {
			actions = append(actions, createActionCommands("all", host)...)

			for _, stack := range stacks {
				actions = append(actions, createActionCommands(stack, host)...)
			}
		}

		actions = append(actions, createActionCommands("all", "all")...)

		slices.SortFunc(actions, func(a, b *action) int {
			return strings.Compare(a.title, b.title)
		})

		actions = slices.CompactFunc(actions, func(a *action, b *action) bool {
			return a.title == b.title
		})

		stacksActions := make(map[string][]*action)
		hostsActions := make(map[string][]*action)

		printActions(actions)

		fmt.Println("===========Stacks actions===========")

		for _, action := range actions {
			for _, stack := range utils.StackList {
				if strings.Contains(action.title, stack) {
					stacksActions[stack] = append(stacksActions[stack], action)
				}
			}

			for _, host := range utils.HostList {
				if strings.Contains(action.title, host) {
					hostsActions[host] = append(hostsActions[host], action)
				}
			}
		}

		for stack, actions := range stacksActions {
			fmt.Println("*****************Stack", stack)
			printActions(actions)
		}

		for host, actions := range hostsActions {
			fmt.Println("*****************Host", host)
			printActions(actions)
		}

	},
}

func init() {
	buildOlivetinCmd.Hidden = true
	rootCmd.AddCommand(buildOlivetinCmd)
}
