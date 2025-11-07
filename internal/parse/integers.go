package parse

import (
	"log"
	"strconv"

	"github.com/insomnes/goaoc/internal/measure"
)

func ToIntegers(input []string) []int {
	defer measure.ExecutionTimeOfParsing("integers", len(input))()
	nums := make([]int, len(input))
	for i, line := range input {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to parse input line %q: %v", line, err)
		}
		nums[i] = n
	}

	return nums
}

func ToUIntegers(input []string) []uint {
	defer measure.ExecutionTimeOfParsing("unsigned integers", len(input))()
	nums := make([]uint, len(input))
	for i, line := range input {
		n, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse input line %q: %v", line, err)
		}
		nums[i] = uint(n)
	}

	return nums
}
