package parse

import (
	"github.com/insomnes/goaoc/internal/measure"
)

func ToRunes(input []string) [][]rune {
	defer measure.ExecutionTimeOfParsing("runes", len(input))()
	runes := make([][]rune, len(input))
	for i, line := range input {
		runes[i] = []rune(line)
	}

	return runes
}
