package utils

import (
	"os"
	"path"
	"slices"
	"strings"
)

type ConfigDir struct {
	envVar     string
	defaultDir string
}

var (
	ComposeDir = ConfigDir{"COMPOSE_DIR", "/workspaces/home-config/compose"}
	EnvDir     = ConfigDir{"ENV_DIR", "/workspaces/home-config/compose"}
	AnsibleDir = ConfigDir{"ANSIBLE_DIR", "/workspaces/home-config/ansible"}
)

var (
	StackList  = getStacks()
	HostList   = getHosts()
	ActionList = getActions()
)

func getDirInEnv(envVariable string, defaultValue string) string {
	envDir := os.Getenv(envVariable)

	if envDir == "" {
		envDir = defaultValue
	}

	return envDir
}

func GetDir(dir ConfigDir) string {
	return getDirInEnv(dir.envVar, dir.defaultDir)
}

func GetFileInDir(dir ConfigDir, file string) string {
	return path.Join(GetDir(dir), file)
}

func getStacks() []string {
	entries, err := os.ReadDir(GetDir(ComposeDir))
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

func getHosts() []string {
	stacks := getStacks()

	hosts := []string{}

	for _, stack := range stacks {
		hosts = append(hosts, GetHostsByStack(stack)...)
	}

	slices.Sort(hosts)

	return slices.Compact(hosts)
}

func GetHostsByStack(stack string) []string {
	entries, err := os.ReadDir(path.Join(GetDir(ComposeDir), stack))

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

func getActions() []string {
	return []string{"up", "down"}
}
