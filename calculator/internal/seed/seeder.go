package seed

import (
	"calculator/internal/model"
	"calculator/internal/repo"
	"calculator/internal/repo/elasticsearch"
	g "calculator/internal/repo/gorm"
	s "calculator/internal/repo/sql"
	"calculator/internal/service"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

type Seeder struct {
	calculatorService *service.CalculatorService
	historyService    *service.HistoryService
	historyIndexer    *elasticsearch.HistoryIndexer
}

func NewSeeder() *Seeder {
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

	return &Seeder{
		calculatorService: svc,
		historyService:    historySvc,
		historyIndexer:    historyIndexer,
	}
}

func (s *Seeder) SeedCalcHistory(total, batchSize int) {
	const workers = 16

	jobs := make(chan int, total)
	docs := make(chan model.CalcHistory, total)

	var wg sync.WaitGroup

	for w := 0; w < workers; w++ {
		wg.Add(1)

		go func(workerID int) {
			defer wg.Done()

			rng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(workerID)*1000))

			for range jobs {
				var record model.CalcHistory
				var err error
				switch pickMode(rng) {
				case model.ExpressionMode:
					record, err = s.seedExpression(rng)
				case model.EquationMode:
					record, err = s.seedEquation(rng)
				case model.LinearSystemMode:
					record, err = s.seedLinearSystem(rng)
				}

				if err != nil {
					log.Printf("[SEED] save DB failed: %v", err)
					continue
				}
				docs <- record
			}
		}(w)
	}

	for i := 0; i < total; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(docs)
	}()

	batch := make([]model.CalcHistory, 0, batchSize)
	indexed := 0

	for doc := range docs {
		batch = append(batch, doc)

		if len(batch) >= batchSize {
			if err := s.historyIndexer.BulkIndex(batch); err != nil {
				log.Printf("bulk index failed: %v", err)
			} else {
				indexed += len(batch)
			}
			batch = make([]model.CalcHistory, 0, batchSize)
		}
	}

	if len(batch) > 0 {
		if err := s.historyIndexer.BulkIndex(batch); err != nil {
			log.Printf("final bulk index failed: %v", err)
		} else {
			indexed += len(batch)
		}
	}

	log.Printf("bulk indexed %d records", indexed)
}

func (s *Seeder) seedExpression(rng *rand.Rand) (model.CalcHistory, error) {
	expr := randomExpression(rng)

	ans, err, durationMs := s.calculatorService.EvalExpression(expr)
	return s.historyService.SaveOnly(model.ExpressionMode, expr, ans, err, durationMs)
}

func (s *Seeder) seedEquation(rng *rand.Rand) (model.CalcHistory, error) {
	degree := 1 + rng.Intn(2)

	var nums []float64
	if degree == 1 {
		nums = randomLinearEquation(rng)
	} else {
		nums = randomQuadraticEquation(rng)
	}

	line := joinFloatSlice(nums)
	ans, err, durationMs := s.calculatorService.SolveEquation(degree, nums)
	return s.historyService.SaveOnly(model.EquationMode, line, ans, err, durationMs)
}

func (s *Seeder) seedLinearSystem(rng *rand.Rand) (model.CalcHistory, error) {
	matrix := randomLinearSystem(rng)

	ans, err, durationMs := s.calculatorService.SolveLinearSystem(matrix)
	return s.historyService.SaveOnly(model.LinearSystemMode, matrix, ans, err, durationMs)
}

func pickMode(rng *rand.Rand) model.Mode {
	x := rng.Intn(100)
	switch {
	case x < 50:
		return model.ExpressionMode
	case x < 80:
		return model.EquationMode
	default:
		return model.LinearSystemMode
	}
}

func joinFloatSlice(nums []float64) string {
	parts := make([]string, len(nums))
	for i, v := range nums {
		if v == float64(int64(v)) {
			parts[i] = fmt.Sprintf("%.0f", v)
		} else {
			parts[i] = fmt.Sprintf("%g", v)
		}
	}
	return strings.Join(parts, " ")
}
