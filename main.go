package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Mayank3299/BlogAggregator/internal/config"
	"github.com/Mayank3299/BlogAggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db         *database.Queries
	configFile *config.Config
}

func main() {
	file, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", file.DbUrl)
	if err != nil {
		log.Fatalf("error opening connection to db: %v", err)
	}

	dbQueries := database.New(db)

	programState := &state{
		configFile: &file,
		db:         dbQueries,
	}

	cmds := commands{
		commandsList: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

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
