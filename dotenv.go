package main

import (
	"errors"
	"os"
	"strings"
)

type DotEnvClient struct {
}

func NewDotEnvClient() *DotEnvClient {
	return &DotEnvClient{}
}

func (c *DotEnvClient) Read(path, env string) (string, error) {
	envs, err := c.readFile(path)
	if err != nil {
		return "", err
	}
	value, err := c.findEnv(envs, env)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (c *DotEnvClient) ReadBulk(path string) (map[string]string, error) {
	envs, err := c.readFile(path)
	if err != nil {
		return nil, err
	}
	return envs, nil
}

func (c *DotEnvClient) Write(path, value string) error {
	return nil
}

func (c *DotEnvClient) WriteBulk(path string, data map[string]string) error {
	return nil
}

func (c *DotEnvClient) findEnv(envs map[string]string, env string) (string, error) {
	value, ok := envs[env]
	if !ok {
		return "", errors.New("env not found")
	}
	return value, nil
}

func (c *DotEnvClient) fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (c *DotEnvClient) readFile(path string) (map[string]string, error) {
	fileExists := c.fileExists(path)
	if !fileExists {
		return nil, errors.New("file does not exist")
	}
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
