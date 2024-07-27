package main

import (
	"os"
	"strings"
)

func ReadDotEnvFile(path string) (map[string]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	envs := make(map[string]string)
	for _, line := range strings.Split(string(bytes), "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		envs[parts[0]] = parts[1]
	}
	return envs, nil
}
