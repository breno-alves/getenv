package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDotEnvClient_findEnv(t *testing.T) {
	c := NewDotEnvClient()
	value, err := c.findEnv("FOO", []byte("FOO=bar\n"))
	if err != nil {
		t.Error(err)
	}
	if value != "bar" {
		t.Errorf("expected bar, got %s", value)
	}
	assert.Equal(t, "bar", value)
}

func TestDotEnvClient_findEnv_envNotFound(t *testing.T) {
	c := NewDotEnvClient()
	_, err := c.findEnv("FOO", []byte("SOMETHINGELSE=bar\n"))
	assert.EqualError(t, err, "env not found")
}
