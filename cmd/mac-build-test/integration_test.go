// +build integration

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestIntegrationAlwaysSucceeds(t *testing.T) {
	fmt.Println("Integration test")
	assert.Check(t, true, "Always succeeds")
}

func TestIntegrationWithEnvVar(t *testing.T) {
	fmt.Println("Integration test with Env var")
	envValue := os.Getenv("ENV_VAR")
	log.Println(strings.Replace(envValue, "2", "#", 1))
	assert.Equal(t, envValue, "42")
}
