package main

import (
	db "match_pool_back/database"
	"match_pool_back/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
)

var optsConfig struct {
	DB string `short:"d" long:"database" description:"Path to the database file" default:"template/template.db"`
}

func main() {

	_, err := flags.ParseArgs(&optsConfig, os.Args)
	if err != nil {
		panic(err)
	}

	err = db.InitDB(optsConfig.DB)
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":5050") // Start the server on port 5050
}
