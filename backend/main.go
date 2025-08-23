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
	"github.com/joho/godotenv"
)

var optsConfig struct {
	DB string `short:"d" long:"database" description:"Path to the database file" default:"database/template/template.db"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	ip := os.Getenv("ALLOWED_IP")
	port := os.Getenv("PORT")
	ginMode := os.Getenv("GIN_MODE")

	gin.SetMode(ginMode)

	_, err = flags.ParseArgs(&optsConfig, os.Args)
	if err != nil {
		panic(err)
	}

	err = db.InitDB(optsConfig.DB)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{fmt.Sprintf("http://%s", ip)},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.SetUpRoutes(r)

	r.Run(fmt.Sprintf("%s:%s", ip, port)) // Start the server on the specified port
}
