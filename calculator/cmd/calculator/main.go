package main

import (
	"calculator/internal/cli"
	"calculator/internal/util"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [gorm|sql]")
		os.Exit(1)
	}

	util.InitLogger()
	app := cli.NewApp()
	app.Run()
}
