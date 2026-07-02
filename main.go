package main

import (
	"log"
	"os"

	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rcopra/gator/internal/config"
	"github.com/rcopra/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	dbQueries := database.New(db)
	programState := state{cfg: &cfg}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	userInput := os.Args
	if len(userInput) < 2 {
		log.Fatal("not enough arguments")
	}
	cmd := command{name: userInput[1], args: userInput[2:]}

	err = cmds.run(&programState, cmd)
	if err != nil {
		log.Fatal(err)
	}

}
