package main

import (
	"fmt"
	"log"
)

const (
	expPart1 = 2
	expPart2 = 1
)

type PassLine struct {
	Letter rune
	Min    int
	Max    int
	Pass   string
}

func parsePassLine(line string) PassLine {
	var pl PassLine
	fmt.Sscanf(line, "%d-%d %c: %s", &pl.Min, &pl.Max, &pl.Letter, &pl.Pass)
	return pl
}

type ParsedInput = []PassLine

func parseInput(lines []string) ParsedInput {
	passLines := make([]PassLine, 0, len(lines))
	for _, line := range lines {
		passLines = append(passLines, parsePassLine(line))
	}
	return passLines
}

func validatePassPart1(pl PassLine) bool {
	cnt := 0
	for _, r := range pl.Pass {
		if r != pl.Letter {
			continue
		}
		cnt++
		if cnt > pl.Max {
			return false
		}
	}
	return cnt >= pl.Min
}

func part1(input ParsedInput) int {
	validCount := 0
	for _, pl := range input {
		if validatePassPart1(pl) {
			validCount++
		}
	}
	return validCount
}

func validatePassPart2(pl PassLine) bool {
	pos1 := pl.Min - 1
	pos2 := pl.Max - 1
	pLen := len(pl.Pass)

	expected := byte(pl.Letter)
	at1, at2 := false, false
	if pos1 >= 0 && pos1 < pLen {
		at1 = pl.Pass[pos1] == expected
	}
	if pos2 >= 0 && pos2 < pLen {
		at2 = pl.Pass[pos2] == expected
	}

	if !at1 && !at2 {
		return false
	}

	return at1 != at2
}

func part2(input ParsedInput) int {
	validCount := 0
	for _, pl := range input {
		if validatePassPart2(pl) {
			validCount++
		}
	}
	return validCount
}

type Solution struct{}

func (s Solution) Part1(input []string) {
	parsed := parseInput(input)
	result := part1(parsed)
	log.Printf("Part 1: %d", result)
}

func (s Solution) Part2(input []string) {
	parsed := parseInput(input)
	result := part2(parsed)
	log.Printf("Part 2: %d", result)
}

func (s Solution) CheckPart1(input []string) error {
	parsed := parseInput(input)
	result := part1(parsed)
	if result != expPart1 {
		return fmt.Errorf("expected %d, got %d", expPart1, result)
	}
	return nil
}

func (s Solution) CheckPart2(input []string) error {
	parsed := parseInput(input)
	result := part2(parsed)
	if result != expPart2 {
		return fmt.Errorf("expected %d, got %d", expPart2, result)
	}
	return nil
}
