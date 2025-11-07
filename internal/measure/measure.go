package measure

import (
	"log"
	"time"
)

func ExecutionTimeOf(name string) func() {
	start := time.Now()
	return func() {
		log.Printf("%s took %v\n", name, time.Since(start))
	}
}

func ExecutionTimeOfParsing(name string, count int) func() {
	start := time.Now()
	return func() {
		log.Printf("Parsing of %s (%d lines) took %v\n", name, count, time.Since(start))
	}
}
