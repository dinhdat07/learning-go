package cli

import (
	"bufio"
	"calculator/internal/repo"
	"calculator/internal/repo/elasticsearch"
	g "calculator/internal/repo/gorm"
	s "calculator/internal/repo/sql"
	"calculator/internal/service"
	"calculator/internal/util"
	"fmt"
	"log"
	"os"
)

type App struct {
	calculatorService *service.CalculatorService
	historyService    *service.HistoryService
	reader            *bufio.Reader
}

func NewApp() *App {
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

	es, err := elasticsearch.Connect()
	if err != nil {
		log.Fatal("Error connecting to ElasticSearch:", err)
	}

	historyIndexer := elasticsearch.NewHistoryIndexer(es)

	svc := service.NewCalculatorService()
	historySvc := service.NewHistoryService(historyRepo, *historyIndexer)

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
