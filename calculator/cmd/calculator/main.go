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
	calculator  *solver.Calculator
	historyRepo *storage.HistoryRepo
	reader      *bufio.Reader
}

func NewApp(repo *storage.HistoryRepo) *App {
	return &App{
		historyRepo: repo,
		calculator:  solver.NewCalculator(),
		reader:      bufio.NewReader(os.Stdin),
	}
}

func (app *App) Run() {
	fmt.Println("CLI CALCULATOR")
	fmt.Println("Type a menu number and press Enter.")

	for {
		opt := utils.ReadInt(app.reader,
			"\nMain Menu\n"+
				"1. Expression\n"+
				"2. Equation\n"+
				"3. Linear System\n"+
				"4. History\n"+
				"5. Exit\n"+
				"> ",
		)

		switch opt {
		case 1:
			app.runExpression()
		case 2:
			app.runEquation()
		case 3:
			app.runLinearSystem()
		case 4:
			// app.runHistory()
		case 5:
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Printf("Invalid option (%d). Please choose 1, 2, 3, or 4.\n", opt)
		}
	}
}

func main() {
	db, err := storage.Connect()
	historyRepo := storage.NewHistoryRepo(db)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
	defer db.Close()

	app := NewApp(historyRepo)
	app.Run()
}
