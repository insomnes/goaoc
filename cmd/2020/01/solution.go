package main

import (
	"fmt"
	"log"

	"github.com/insomnes/goaoc/internal/parse"
)

const (
	expPart1 = 514579
	expPart2 = 241861950 // No expected value provided yet
	toMeet   = 2020
)

type ParsedInput = []int

func recFind(input ParsedInput, target int, depth int) (int, bool) {
	if depth == 0 {
		return -1, false
	}

	for i, n := range input {
		newTarget := target - n
		if newTarget < 0 {
			continue
		}
		if newTarget == 0 && depth == 1 {
			return n, true
		}
		subResult, found := recFind(input[i+1:], newTarget, depth-1)
		if found {
			return n * subResult, true
		}
	}
	return -1, false
}

func part1(input ParsedInput) int {
	result, found := recFind(input, toMeet, 2)
	if found {
		return result
	}
	log.Fatalf("No solution found")
	return -1
}

func part2(input ParsedInput) int {
	if len(input) < 3 {
		log.Fatalf("Input too short")
	}
	result, found := recFind(input, toMeet, 3)
	if !found {
		log.Fatalf("No solution found")
	}
	return result
}

type Solution struct{}

func (s Solution) Part1(input []string) {
	parsed := parse.ToIntegers(input)
	result := part1(parsed)
	log.Printf("Part 1: %d", result)
}

func (s Solution) Part2(input []string) {
	parsed := parse.ToIntegers(input)
	result := part2(parsed)
	log.Printf("Part 2: %d", result)
}

func (s Solution) CheckPart1(input []string) error {
	parsed := parse.ToIntegers(input)
	result := part1(parsed)
	if result != expPart1 {
		return fmt.Errorf("expected %d, got %d", expPart1, result)
	}
	return nil
}

func (s Solution) CheckPart2(input []string) error {
	parsed := parse.ToIntegers(input)
	result := part2(parsed)
	if result != expPart2 {
		return fmt.Errorf("expected %d, got %d", expPart2, result)
	}
	return nil
}
