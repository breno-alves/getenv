package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDotEnvClient_findEnv(t *testing.T) {
	c := NewDotEnvClient()
	value, err := c.findEnv(map[string]string{"FOO": "bar"}, "FOO")
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
	_, err := c.findEnv(map[string]string{"SOMETHINGELSE": "bar"}, "FOO")
	assert.EqualError(t, err, "env not found")
}
