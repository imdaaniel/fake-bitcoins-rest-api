package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"

	"github.com/imdaaniel/bitcoins-rest-api/api/models"
)

type Server struct {
	DB		*gorm.DB
	Router	*mux.Router
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	server.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Order{}) //migration of database

	server.Router = mux.NewRouter()

	server.InitializeRoutes()
}

func (server *Server) Run(address string) {
	fmt.Println("Listening on port %s", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}