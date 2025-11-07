package main

import (
	"fmt"
	"log"

	"github.com/insomnes/goaoc/internal/measure"
	"github.com/insomnes/goaoc/internal/parse"
)

const (
	expPart1 = 7
	expPart2 = 336
)

type ParsedInput = [][]rune

const (
	tree = '#'
)

type Slope struct {
	Dx int
	Dy int
}

var (
	p1slope  = Slope{Dx: 3, Dy: 1}
	p2slopes = []Slope{
		{Dx: 1, Dy: 1},
		{Dx: 3, Dy: 1},
		{Dx: 5, Dy: 1},
		{Dx: 7, Dy: 1},
		{Dx: 1, Dy: 2},
	}
)

func countTrees(slopeMap ParsedInput, slope Slope) int {
	metTrees := 0
	posX, posY := 0, 0
	dx, dy := slope.Dx, slope.Dy

	rowLen := len(slopeMap[0])

	for {
		posX = (posX + dx) % rowLen
		posY += dy
		if posY >= len(slopeMap) {
			break
		}
		c := slopeMap[posY][posX]
		if c == tree {
			metTrees += 1
		}
	}
	return metTrees
}

func part1(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 1")()
	return countTrees(input, p1slope)
}

func part2(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 2")()
	answer := 1
	for _, slope := range p2slopes {
		answer *= countTrees(input, slope)
	}
	return answer
}

type Solution struct{}

func (s Solution) Part1(input []string) {
	parsed := parse.ToRunes(input)
	result := part1(parsed)
	log.Printf("Part 1 answer: %d\n", result)
}

func (s Solution) Part2(input []string) {
	parsed := parse.ToRunes(input)
	result := part2(parsed)
	log.Printf("Part 2 answer: %d\n", result)
}

func (s Solution) CheckPart1(input []string) error {
	parsed := parse.ToRunes(input)
	result := part1(parsed)
	log.Printf("Checking Part 1: got %d, expected %d\n", result, expPart1)
	if result != expPart1 {
		return fmt.Errorf("expected %d, got %d", expPart1, result)
	}
	return nil
}

func (s Solution) CheckPart2(input []string) error {
	parsed := parse.ToRunes(input)
	result := part2(parsed)
	log.Printf("Checking Part 2: got %d, expected %d\n", result, expPart2)
	if result != expPart2 {
		return fmt.Errorf("expected %d, got %d", expPart2, result)
	}
	return nil
}
