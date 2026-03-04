package main

import (
	"calculator/internal/utils"
	"fmt"
)

func (app *App) runHistory() {
	for {
		opt := utils.ReadInt(app.reader,
			"\nHistory Menu\n"+
				"1. List history\n"+
				"2. View history by ID\n"+
				"3. Delete history by ID\n"+
				"4. Clear all history\n"+
				"5. Back\n"+
				"> ",
		)

		switch opt {
		case 1:
			app.listHistory()
		case 2:
			app.viewHistory()
		case 3:
			app.deleteHistory()
		case 4:
			app.clearHistory()
		case 5:
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func (app *App) clearHistory() {
	panic("unimplemented")
}

func (app *App) deleteHistory() {
	panic("unimplemented")
}

func (app *App) viewHistory() {
	panic("unimplemented")
}

func (app *App) listHistory() {
	panic("unimplemented")
}
