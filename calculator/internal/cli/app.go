package cli

import (
	"bufio"
	"calculator/internal/repo"
	"calculator/internal/repo/elasticsearch"
	"calculator/internal/service"
	util "calculator/internal/util"
	"fmt"
	"os"
)

type App struct {
	calculatorService *service.CalculatorService
	historyService    *service.HistoryService
	reader            *bufio.Reader
}

func NewApp(repo repo.HistoryRepo, indexer elasticsearch.HistoryIndexer) *App {
	// repo
	// service
	svc := service.NewCalculatorService()
	historySvc := service.NewHistoryService(repo, indexer)

	return &App{
		calculatorService: svc,
		historyService:    historySvc,
		reader:            bufio.NewReader(os.Stdin),
	}
}

func (app *App) Run() {
	fmt.Println("CLI CALCULATOR")
	fmt.Println("Type a menu number and press Enter.")

	for {
		opt := util.ReadInt(app.reader,
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
			app.runHistory()
		case 5:
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Printf("Invalid option (%d). Please choose 1, 2, 3, or 4.\n", opt)
		}
	}
}
