package main

import (
	"fmt"
	"log"
	"math/bits"

	"github.com/insomnes/goaoc/internal/dlog"
	"github.com/insomnes/goaoc/internal/dsa/set"
	"github.com/insomnes/goaoc/internal/measure"
	"github.com/insomnes/goaoc/internal/parse"
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
	result := make([]Group, len(lines))
	var current Group

	for _, line := range lines {
		if line == "" {
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

func findGroupAnswersAnyBitset(group Group) uint {
	var allAnswers uint = 0
	for _, person := range group.persons {
		if allAnswers == (1<<maxAnswers)-1 {
			dlog.Debugf("All answers found, breaking early")
			break
		}
		persAnswers := parse.AlphaStrToBitMask(person)
		allAnswers |= persAnswers
	}
	return allAnswers
}

func part1(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 1")()

	result := 0

	for i, group := range input {
		dlog.Debugf("Checking group %d: %v", i, group)
		groupAnswers := findGroupAnswersAnyBitset(group)
		result += bits.OnesCount(groupAnswers)
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

func FindGroupAnswersAllBitset(group Group) uint {
	minPerson := group.minPerson
	var commonAnswers uint = 0
	commonAnswers = parse.AlphaStrToBitMask(minPerson)

	for _, person := range group.persons {
		if commonAnswers == 0 {
			dlog.Debugf("No common answers left, breaking early")
			break
		}
		personAnswers := parse.AlphaStrToBitMask(person)
		commonAnswers &= personAnswers
	}

	return commonAnswers
}

func part2(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 2")()
	result := 0

	for i, group := range input {
		dlog.Debugf("Checking group %d: %v", i, group)
		groupAnswers := FindGroupAnswersAllBitset(group)
		result += bits.OnesCount(groupAnswers)
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
