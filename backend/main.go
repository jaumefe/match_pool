package main

import (
	"fmt"
	db "match_pool_back/database"
	"match_pool_back/routes"
	"os"
	"time"

	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetUpRoutes(r)

	port := fmt.Sprintf(":%s", optsConfig.Port)
	r.Run(port) // Start the server on the specified port
}
