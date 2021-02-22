package main

import (
	"os"
	"strconv"
)

var app App

func main() {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		panic("Invalid Database port")
	} else {
		app.init(
			os.Getenv("DATABASE_HOST"),
			port,
			os.Getenv("DATABASE_USERNAME"),
			os.Getenv("DATABASE_PASSWORD"))
		app.Run(":8090")
	}

}
