package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	if os.Getenv("TEST_EXIT_CODE") == "true" {
		main()
		return
	}

	t.Run("exit code 0", func(t *testing.T) {
		cmd := exec.Command(os.Args[0], "./tests/step1/valid.json", "-test.run=TestMain_exit_code_0")
		cmd.Env = append(os.Environ(), "TEST_EXIT_CODE=true")

		err := cmd.Run()

		assert.NoError(t, err)
	})

	t.Run("exit code 1", func(t *testing.T) {
		cmd := exec.Command(os.Args[0], "./tests/step1/invalid.json", "-test.run=TestMain_exit_code_1")
		cmd.Env = append(os.Environ(), "TEST_EXIT_CODE=true")

		err := cmd.Run()
		e, ok := err.(*exec.ExitError)

		assert.True(t, ok)
		assert.False(t, e.Success())
	})
}
