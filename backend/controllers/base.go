package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	//_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	//_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	//"github.com/victorsteven/fullstack/api/models"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
