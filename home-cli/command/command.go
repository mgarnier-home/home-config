package command

import (
	"fmt"
	"mgarnier11/home-cli/utils"
	"os/exec"
	"slices"
)

type Command struct {
	action string
	stack  string
	host   string
}

func getCommandsByStack(stack string, host string, action string) []*Command {
	commands := []*Command{}

	hosts := utils.GetHostsByStack(stack)

	for _, h := range hosts {
		if host == "all" || h == host {
			commands = append(commands, &Command{action, stack, h})
		}
	}

	return commands
}

func getCommandsToExecute(stack string, host string, action string) []*Command {

	commands := []*Command{}

	if stack == "all" {
		stacks := utils.GetStacks()

		for _, stack := range stacks {
			commands = append(commands, getCommandsByStack(stack, host, action)...)
		}
	} else {
		commands = append(commands, getCommandsByStack(stack, host, action)...)
	}

	return commands
}

func ExecCommand(stack string, host string, action string, args []string) {
	fmt.Printf("Running %s on stack : %s host : %s\n", action, stack, host)

	commands := getCommandsToExecute(stack, host, action)

	commandsPerHost := map[string][]*Command{}

	for _, command := range commands {
		commandsPerHost[command.host] = append(commandsPerHost[command.host], command)
	}

	if slices.Contains(args, "parallel") {
		runCommandsInParallel(commandsPerHost)
	} else {
		runCommandsInSequence(commandsPerHost)
	}
}

func runCommandsInParallel(commandsPerHost map[string][]*Command) {

}

func runCommandsInSequence(commandsPerHost map[string][]*Command) {
	for _, commands := range commandsPerHost {
		selectHost(commands[0].host)

		for _, command := range commands {
			execCommand(command)
		}

		selectHost("default")
	}

}

func selectHost(host string) {
	out, err := exec.Command("docker", "context", "use", host).CombinedOutput()
	if err != nil {
		fmt.Println("Error selecting host", err)
		panic(err)
	}

	fmt.Println(string(out))

}

func execCommand(command *Command) {
	// fmt.Printf("Executing command %s on host %s stack %s\n", command.action, command.host, command.stack)

	var commandArgs = []string{
		"docker",
		"compose",
		"--env-file compose/env.env",
		"--env-file compose/ports.env",
		"-f compose/volumes.yml",
		"-f compose/" + command.stack + "/" + command.host + "." + command.stack + ".yml",
	}

	if command.action == "up" {
		commandArgs = append(commandArgs, "up", "-d")
	} else if command.action == "down" {
		commandArgs = append(commandArgs, "down", "-v")
	} else if command.action == "pull" {

	} else {
		fmt.Println("Command not found")
	}

	fmt.Println(commandArgs)

	out, err := exec.Command(commandArgs[0], commandArgs[1:]...).CombinedOutput()

	if err != nil {
		fmt.Println("error executing command ", &command, err)
	}

	fmt.Println(string(out))
}
