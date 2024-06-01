package utils

import (
	"os"
	"path"
	"slices"
	"strings"
)

func getDirInEnv(envVariable string, defaultValue string) string {
	envDir := os.Getenv(envVariable)

	if envDir == "" {
		envDir = defaultValue
	}

	return envDir
}

func getComposeDir() string {
	return getDirInEnv("COMPOSE_DIR", "/workspaces/home-config/compose")
}

func getEnvDir() string {
	return getDirInEnv("ENV_DIR", "/workspaces/home-config/compose")
}

func GetFileInComposeDir(file string) string {
	return path.Join(getComposeDir(), file)
}

func GetFileInEnvDir(file string) string {
	return path.Join(getEnvDir(), file)
}

func GetStacks() []string {
	entries, err := os.ReadDir(getComposeDir())
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
	entries, err := os.ReadDir(path.Join(getComposeDir(), stack))

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
	return []string{"up", "down" /*, "pull"*/}
}
