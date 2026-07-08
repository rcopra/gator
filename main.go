package main

import (
	"fmt"
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
		log.Fatal("Error reading config")
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal("Error viewing database")
	}
	dbQueries := database.New(db)
	programState := &state{db: dbQueries, cfg: &cfg}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmds.register("feeds", handlerGetFeeds)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	userInput := os.Args
	if len(userInput) < 2 {
		log.Fatal("not enough arguments")
	}
	cmd := command{Name: userInput[1], Args: userInput[2:]}

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Completed actions successfully...terminating")

}
