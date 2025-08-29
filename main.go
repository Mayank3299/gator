package main

import (
	"log"
	"os"

	"github.com/Mayank3299/BlogAggregator/internal/config"
)

type state struct {
	configFile *config.Config
}

func main() {
	file, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		configFile: &file,
	}

	cmds := commands{
		commandsList: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
