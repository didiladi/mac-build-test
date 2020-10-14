// +build cleanup

package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestCleanupAlwaysSucceeds(t *testing.T) {
	fmt.Println("Cleanup")
	assert.Check(t, true, "Always succeeds")
}