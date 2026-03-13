package cli

import (
	"calculator/internal/util"
	"fmt"
	"log"
	"strings"
)

func (app *App) runHistory() {
	for {
		opt := util.ReadInt(app.reader,
			"\nHistory Menu\n"+
				"1. List history\n"+
				"2. Update history note\n"+
				"3. Delete history by ID\n"+
				"4. Clear all history\n"+
				"5. Back\n"+
				"> ",
		)

		switch opt {
		case 1:
			app.listHistory()
		case 2:
			app.updateHistoryNote()
		case 3:
			app.deleteHistory()
		case 4:
			app.clearHistory()
		case 5:
			fmt.Println("Returning to main menu.")
			return
		default:
			fmt.Println("Invalid option. Please choose 1, 2, 3, 4, or 5.")
		}
	}
}

func (app *App) listHistory() {
	fmt.Println("\nHistory List")
	fmt.Println("Enter number of records to show (<= 0 to show all).")
	limit := util.ReadInt(app.reader, "> ")

	list, err := app.historyService.List(limit)
	if err != nil {
		log.Printf("Internal error with database: %v\n", err)
		return
	}

	if len(list) == 0 {
		fmt.Println("No history records found.")
		return
	}

	for i := range list {
		fmt.Println(list[i])
	}
}

func (app *App) updateHistoryNote() {
	fmt.Println("\nUpdate History Note")
	id := util.ReadInt(app.reader, "Enter history record ID:\n> ")

	fmt.Println("Enter note (empty to clear note):")
	note := util.ReadLine(app.reader, "> ")

	if err := app.historyService.UpdateNote(int64(id), note); err != nil {
		log.Printf("Internal error with database: %v\n", err)
		return
	}

	fmt.Println("Note updated.")
}

func (app *App) deleteHistory() {
	fmt.Println("\nDelete History Record")
	id := util.ReadInt(app.reader, "Enter history record ID:\n> ")

	confirm := strings.ToUpper(util.ReadLine(app.reader, "Are you sure you want to delete this record? (Y/N)\n> "))
	if confirm != "Y" {
		fmt.Println("Cancelled.")
		return
	}

	if err := app.historyService.Delete(id); err != nil {
		log.Printf("Internal error with database: %v\n", err)
		return
	}

	fmt.Println("History record deleted.")
}

func (app *App) clearHistory() {
	fmt.Println("\nClear All History")
	confirm := strings.ToUpper(util.ReadLine(app.reader, "This will delete ALL history records. Continue? (Y/N)\n> "))
	if confirm != "Y" {
		fmt.Println("Cancelled.")
		return
	}

	if err := app.historyService.Clear(); err != nil {
		log.Printf("Internal error with database: %v\n", err)
		return
	}

	fmt.Println("All history records deleted.")
}
