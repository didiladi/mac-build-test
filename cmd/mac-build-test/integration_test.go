// +build integration

package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestIntegrationAlwaysSucceeds(t *testing.T) {
	fmt.Println("Integration test")
	assert.Check(t, true, "Always succeeds")
}