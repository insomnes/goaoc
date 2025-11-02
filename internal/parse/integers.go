package parse

import (
	"log"
	"strconv"
	"time"
)

func ToIntegers(input []string) []int {
	start := time.Now()
	nums := make([]int, len(input))
	for i, line := range input {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to parse input line %q: %v", line, err)
		}
		nums[i] = n
	}

	log.Printf("Parsed %d integers in %s", len(nums), time.Since(start))
	return nums
}

func ToUIntegers(input []string) []uint {
	start := time.Now()
	nums := make([]uint, len(input))
	for i, line := range input {
		n, err := strconv.ParseUint(line, 10, 64)
		if err != nil {
			log.Fatalf("Failed to parse input line %q: %v", line, err)
		}
		nums[i] = uint(n)
	}

	log.Printf("Parsed %d unsigned integers in %s", len(nums), time.Since(start))

	return nums
}
