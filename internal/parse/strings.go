package parse

import (
	"log"
	"time"
)

func ToRunes(input []string) [][]rune {
	start := time.Now()
	runes := make([][]rune, len(input))
	for i, line := range input {
		runes[i] = []rune(line)
	}

	log.Printf("Parsed %d lines of runes in %s", len(runes), time.Since(start))
	return runes
}
