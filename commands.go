package main

import (
	"errors"
	"fmt"

	"github.com/Mayank3299/BlogAggregator/internal/database"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	commandsList map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commandsList[cmd.Name]
	if !ok {
		return errors.New("command doesn't exist")
	}

	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandsList[name] = f
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
