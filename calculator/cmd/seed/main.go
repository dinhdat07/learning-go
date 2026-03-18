package main

import (
	"calculator/internal/seed"
	"calculator/internal/util"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go [gorm|sql] [total] [batchSize]")
		os.Exit(1)
	}

	util.InitLogger()

	total, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("invalid total: %v", err)
	}

	batchSize, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("invalid total: %v", err)
	}

	seeder := seed.NewSeeder()
	if seeder == nil {
		log.Fatal("failed to initialize seeder")
	}

	seeder.SeedCalcHistory(total, batchSize)

	log.Printf("seeded %d calc_history records\n", total)
}
