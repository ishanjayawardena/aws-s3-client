package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Displays the input message in the OS pager
func displayInPager(msg string) error {
	pager := exec.Command(os.ExpandEnv("$PAGER"))
	pager.Stdin = strings.NewReader(msg)
	pager.Stdout = os.Stdout
	err := pager.Run()

	if err != nil {
		return fmt.Errorf("Pager error: %v", err)
	}

	return nil
}
