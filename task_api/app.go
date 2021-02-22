package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) init(host string, port int, username string, password string) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=task_manager sslmode=disable TimeZone=America/Lima",
		host, port, username, password)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&Task{})
	a.DB = database
	a.Router = mux.NewRouter().StrictSlash(true)
}

func (app *App) Run(addr string) {
	app.Router.HandleFunc("/tasks", findAll).Methods("GET")
	app.Router.HandleFunc("/tasks/{id}", findByID).Methods("GET")
	app.Router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	app.Router.HandleFunc("/tasks", createTask).Methods("POST")

	if err := http.ListenAndServe(addr, app.Router); err != nil {
		log.Fatal(err)
	} else {
		log.Output(1, "Successfully started http server")
	}
}
