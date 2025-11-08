package main

import (
	"fmt"
	"log"

	"github.com/insomnes/goaoc/internal/dlog"
	"github.com/insomnes/goaoc/internal/dsa/set"
	"github.com/insomnes/goaoc/internal/measure"
)

const (
	expPart1 = 11
	expPart2 = 6
)

type Group struct {
	persons   []string
	minPerson string
}

func (g *Group) Len() int {
	return len(g.persons)
}

func (g *Group) AddPerson(p string) {
	g.persons = append(g.persons, p)
	if g.minPerson == "" || len(p) < len(g.minPerson) {
		g.minPerson = p
	}
}

const maxAnswers = 26

type ParsedInput = []Group

func parseInput(lines []string) ParsedInput {
	defer measure.ExecutionTimeOf("Parse Input")()
	result := make([]Group, 0)
	var current Group

	for _, line := range lines {
		if line == "" {
			if current.Len() == 0 {
				log.Fatal("Empty group")
			}
			result = append(result, current)
			current = Group{}
			continue
		}

		current.AddPerson(line)
	}

	if current.Len() > 0 {
		result = append(result, current)
	}

	return result
}

func findGroupAnswersAny(group Group) *set.Set[rune] {
	allAnswers := set.NewSetCap[rune](maxAnswers)

	for _, person := range group.persons {
		if allAnswers.Size() == maxAnswers {
			break
		}
		for _, ans := range person {
			allAnswers.Add(ans)
		}
	}

	return allAnswers
}

func part1(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 1")()

	result := 0

	for i, group := range input {
		dlog.Debugf("Checking group %d: %v", i, group)
		result += findGroupAnswersAny(group).Size()
	}
	return result
}

func FindGroupAnswersAll(group Group) *set.Set[rune] {
	minPerson := group.minPerson
	commonAnswers := set.NewSetCap[rune](len(minPerson))

	for _, ans := range minPerson {
		commonAnswers.Add(ans)
	}

	for _, person := range group.persons {
		if commonAnswers.IsEmpty() {
			break
		}
		personAnswers := set.NewSetCap[rune](len(person))
		for _, ans := range person {
			personAnswers.Add(ans)
		}
		commonAnswers = commonAnswers.Intersection(personAnswers)

	}
	return commonAnswers
}

func part2(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 2")()
	result := 0

	for i, group := range input {
		dlog.Debugf("Checking group %d: %v", i, group)
		result += FindGroupAnswersAll(group).Size()
	}
	return result
}

type Solution struct{}

func (s Solution) Part1(input []string) {
	parsed := parseInput(input)
	result := part1(parsed)
	log.Printf("Part 1 answer: %d\n", result)
}

func (s Solution) Part2(input []string) {
	parsed := parseInput(input)
	result := part2(parsed)
	log.Printf("Part 2 answer: %d\n", result)
}

func (s Solution) CheckPart1(input []string) error {
	parsed := parseInput(input)

	result := part1(parsed)
	log.Printf("Checking Part 1: got %d, expected %d\n", result, expPart1)
	if result != expPart1 {
		return fmt.Errorf("expected %d, got %d", expPart1, result)
	}
	return nil
}

func (s Solution) CheckPart2(input []string) error {
	parsed := parseInput(input)
	result := part2(parsed)
	log.Printf("Checking Part 2: got %d, expected %d\n", result, expPart2)
	if result != expPart2 {
		return fmt.Errorf("expected %d, got %d", expPart2, result)
	}
	return nil
}
