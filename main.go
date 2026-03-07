package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/gutek00714/Blog-Aggregator---Boot.dev/internal/config"
	"github.com/gutek00714/Blog-Aggregator---Boot.dev/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	//open a connection to database
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dbQueries := database.New(db)

	// store the config in a new instance of the state struct
	programState := &state{
		Config: &cfg,
		db:     dbQueries,
	}

	// initialize the map
	cmds := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	// register the command
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerResetDatabase)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)

	// check os.Args length (if enough arguments were provided)
	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	// build the command from os.Args
	cmd := command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	// run the command
	err = cmds.run(programState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
