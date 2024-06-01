package command

import (
	"bufio"
	"fmt"
	"io"
	"mgarnier11/home-cli/utils"
	"os/exec"
	"slices"
	"sync"

	"github.com/fatih/color"
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

	for _, commands := range commandsPerHost {
		selectHost(commands[0].host)

		if slices.Contains(args, "parallel") {
			var wg sync.WaitGroup

			for _, command := range commands {
				wg.Add(1)

				go (func(command *Command, wg *sync.WaitGroup) {
					defer wg.Done()

					execCommand(command)
				})(command, &wg)
			}

			wg.Wait()
		} else {
			for _, command := range commands {
				execCommand(command)
			}
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

func print(command *Command, std io.ReadCloser) {
	scanner := bufio.NewScanner(std)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()

		color.Yellow(fmt.Sprintf("%s %s %s", command.host, command.stack, text))
	}
}

func execCommand(command *Command) {
	var commandArgs = []string{
		"docker",
		"compose",
		"--env-file",
		utils.GetFileInEnvDir("env.env"),
		"--env-file",
		utils.GetFileInEnvDir("ports.env"),
		"-f",
		utils.GetFileInComposeDir("volumes.yml"),
		"-f",
		utils.GetFileInComposeDir(command.stack + "/" + command.host + "." + command.stack + ".yml"),
	}

	if command.action == "up" {
		commandArgs = append(commandArgs, "up", "-d", "--pull", "always")
	} else if command.action == "down" {
		commandArgs = append(commandArgs, "down", "-v")
	} else {
		fmt.Println("Command not found")
	}

	color.Blue(fmt.Sprint(commandArgs))

	cmd := exec.Command(commandArgs[0], commandArgs[1:]...)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	go print(command, stdout)
	go print(command, stderr)

	err := cmd.Run()

	if err != nil {
		color.Red(fmt.Sprintf("%s %s Error executing command %s", command.host, command.stack, err))
	} else {
		color.Green(fmt.Sprintf("%s %s Successfully executed command", command.host, command.stack))
	}
}
