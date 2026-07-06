package main

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command")
	}
	err := handler(s, cmd)
	return err
}

func (c *commands) register(name string, f func(s *state, cmd command) error) {
	c.registeredCommands[name] = f
}
