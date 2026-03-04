package main

import (
	"bufio"
	"calculator/internal/solver"
	"calculator/internal/storage"
	"calculator/internal/utils"
	"fmt"
	"log"
	"os"
)

type App struct {
	calculator *solver.Calculator
	db *storage.DB
	reader *bufio.Reader
}


func NewApp(db *sql.DB) *App {
	return &App{
		db:         db,
		calculator: solver.NewCalculator(),
		reader:     bufio.NewReader(os.Stdin),
	}
}


func (app *App) Run() {
	fmt.Println("CLI CALCULATOR")
	fmt.Println("Type a menu number and press Enter.")

	for {
		opt := utils.ReadInt(reader,
			"\nMain Menu\n"+
				"1. Expression\n"+
				"2. Equation\n"+
				"3. Linear System\n"+
				"4. Exit\n"+
				"> ",
		)

		switch opt {
		case 1:
			a.runExpression()
		case 2:
			a.runEquation()
		case 3:
			a.runLinearSystem()
		case 4:
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Printf("Invalid option (%d). Please choose 1, 2, 3, or 4.\n", opt)
		}
	}
}



func main() {
	db, err := storage.Connect()
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
	defer db.Close()

	app := NewApp(db)
	app.Run()
}