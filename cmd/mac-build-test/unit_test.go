// +build unit

package main

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestUnitAlwaysSucceeds(t *testing.T) {
	fmt.Println("Unit test")
	assert.Check(t, true, "Always succeeds")
}