package utils

import (
	"errors"
	"os"
	"slices"
	"strings"
)

func StackHostExists(stack string, host string) bool {
	_, err := os.Stat("./compose/" + stack + "/" + host + "." + stack + ".yml")

	if err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		panic(err)
	}
}

func GetStacks() []string {
	entries, err := os.ReadDir("./compose")
	if err != nil {
		return []string{}
	}

	stacks := []string{}

	for _, entry := range entries {
		if entry.IsDir() {
			stacks = append(stacks, entry.Name())
		}
	}

	slices.Sort(stacks)

	return stacks
}

func GetHosts() []string {
	stacks := GetStacks()

	hosts := []string{}

	for _, stack := range stacks {
		hosts = append(hosts, GetHostsByStack(stack)...)
	}

	slices.Sort(hosts)

	return slices.Compact(hosts)
}

func GetHostsByStack(stack string) []string {
	entries, err := os.ReadDir("./compose/" + stack)

	if err != nil {
		return []string{}
	}

	hosts := []string{}

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), "."+stack+".yml") {
			parts := strings.Split(entry.Name(), ".")

			hosts = append(hosts, parts[0])
		}
	}

	return hosts
}

func GetActions() []string {
	return []string{"up", "down", "pull"}
}
