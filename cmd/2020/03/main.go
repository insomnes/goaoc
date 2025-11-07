package main

import (
	"log"
	"os"

	"github.com/insomnes/goaoc/internal/run"
)

const (
	year = 2020
	day  = 3
)

func main() {
	// Check If anything is provided, then run in solve mode
	if len(os.Args) > 1 {
		if err := run.Solve(year, day, Solution{}); err != nil {
			log.Fatalf("Error: %v", err)
		}
	} else {
		if err := run.Check(year, day, Solution{}); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
