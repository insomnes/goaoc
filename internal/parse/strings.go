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

func AlphaRuneToBit(r rune) uint {
	return uint(1) << (r - 'a')
}

func AlphaStrToBitMask(answers string) uint {
	var bits uint = 0
	for _, ans := range answers {
		bit := AlphaRuneToBit(ans)
		bits |= bit
	}
	return bits
}
