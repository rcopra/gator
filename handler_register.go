package main

import (
	"fmt"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("invalid input missing command")
	}
	userName := cmd.args[0]
}
