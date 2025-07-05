package main

import (
	"fmt"
	db "match_pool_back/database"
	"match_pool_back/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
)

var optsConfig struct {
	DB   string `short:"d" long:"database" description:"Path to the database file" default:"database/template/template.db"`
	Port string `short:"p" long:"port" description:"Port to run the server on" default:"8080"`
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
	routes.SetUpRoutes(r)

	port := fmt.Sprintf(":%s", optsConfig.Port)
	r.Run(port) // Start the server on the specified port
}
