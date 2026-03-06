package main

import (
	"calculator/internal/cli"
	"calculator/internal/repo"
	g "calculator/internal/repo/gorm"
	s "calculator/internal/repo/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [gorm|sql]")
		os.Exit(1)
	}

	var historyRepo repo.HistoryRepo
	mode := os.Args[1]
	switch mode {
	case "gorm":
		db, err := g.Connect()
		if err != nil {
			log.Fatal("Error connecting to DB:", err)
		}
		historyRepo = g.NewHistoryRepo(db)

	case "sql":
		db, err := s.Connect()
		if err != nil {
			log.Fatal("Error connecting to DB:", err)
		}
		defer db.Close()
		historyRepo = s.NewHistoryRepo(db)
	}

	app := cli.NewApp(historyRepo)

	app.Run()
}
