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

// MustBeInt generic integer parser that fatals on error.
func MustBeInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64](s string) T {
	// Parse the string as an integer.
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse integer from %q: %v", s, err)
	}
	// Check for overflow.
	var zero T
	if n < int64(zero) || n > int64(^zero) {
		log.Fatalf("Integer overflow when parsing %q into type %T", s, zero)
	}
	return T(n)
}

// MustBeUint generic unsigned integer parser that fatals on error.
func MustBeUint[T ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](s string) T {
	// Parse the string as an unsigned integer.
	n, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse unsigned integer from %q: %v", s, err)
	}
	// Check for overflow.
	var zero T
	if n > uint64(^zero) {
		log.Fatalf("Unsigned integer overflow when parsing %q into type %T", s, zero)
	}
	return T(n)
}
