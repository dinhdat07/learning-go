package main

import (
	"calculator/internal/cli"
	"calculator/internal/db"
	"log"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
	defer db.Close()
	app := cli.NewApp(db)
	app.Run()
}
