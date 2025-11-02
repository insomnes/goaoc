package main

import (
	"log"
	"os"

	"github.com/insomnes/goaoc/internal/run"
)

const (
	year = 2020
	day  = 1
)

func main() {
	// Check If anything is provided, than run in solve mode
	if len(os.Args) > 1 {
		err := run.Solve(year, day, Solution{})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	} else {
		err := run.Check(year, day, Solution{})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}
}
