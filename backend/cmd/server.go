package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/silvano-bergamasco/business6sense/backend/controllers"
	//"github.com/victorsteven/fullstack/api/controllers"
	//"github.com/victorsteven/fullstack/api/seed"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("MYSQL_DRIVER"), os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))

	//seed.Load(server.DB)

	server.Run(":8090")

}
