package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/insomnes/goaoc/internal/dlog"
	"github.com/insomnes/goaoc/internal/measure"
)

const (
	expPart1 = 820
	expPart2 = -9999 // No expected value available
)

var (
	rowReplacer = strings.NewReplacer("F", "0", "B", "1")
	colReplacer = strings.NewReplacer("L", "0", "R", "1")
)

type ParsedInput = []string

type Seat struct {
	Row    int
	Column int
	ID     int
}

func parseSeat(seatCode string) Seat {
	if len(seatCode) != 10 {
		log.Fatalf("Invalid seat code length: %s", seatCode)
	}
	rowPart := seatCode[:7]
	colPart := seatCode[7:]

	rowBin := rowReplacer.Replace(rowPart)
	row, err := strconv.ParseInt(rowBin, 2, 64)
	if err != nil {
		log.Fatalf("Failed to parse row from %s: %v", rowBin, err)
	}

	colBin := colReplacer.Replace(colPart)
	col, err := strconv.ParseInt(colBin, 2, 64)
	if err != nil {
		log.Fatalf("Failed to parse column from %s: %v", colBin, err)
	}
	finalRow := int(row)
	finalCol := int(col)
	seatID := finalRow*8 + finalCol

	return Seat{
		Row:    finalRow,
		Column: finalCol,
		ID:     seatID,
	}
}

func part1(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 1")()
	maxID := 0
	for _, seatCode := range input {
		seat := parseSeat(seatCode)
		dlog.Debugf("Seat: %+v", seat)
		if seat.ID > maxID {
			maxID = seat.ID
		}
	}
	return maxID
}

func part2(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 2")()
	allSeats := make([]bool, 1024)
	minId, maxId := 1024, 0
	sum := 0
	for _, seatCode := range input {
		seat := parseSeat(seatCode)
		if seat.ID < minId {
			minId = seat.ID
		}
		if seat.ID > maxId {
			maxId = seat.ID
		}
		sum += seat.ID
		allSeats[seat.ID] = true
	}
	dlog.Debugf("PART 2: Seat ID range: %d - %d, sum %d\n", minId, maxId, sum)

	expectedSum := (maxId - minId + 1) * (minId + maxId) / 2
	missingId := expectedSum - sum
	if missingId >= len(allSeats) || missingId < 0 {
		log.Printf(
			"[CRITICAL] Calculated missing seat ID %d is out of bounds (checking?!)",
			missingId,
		)
		return -1
	}

	log.Printf("PART 2: Calculated missing seat ID: %d", missingId)
	if allSeats[missingId-1] && allSeats[missingId+1] {
		log.Printf("PART 2: Found my seat ID: %d", missingId)
	} else {
		log.Printf("PART 2: Seat ID %d is not my seat (neighbors not occupied)", missingId)
	}
	return missingId
}

type Solution struct{}

func (s Solution) Part1(input []string) {
	parsed := input
	result := part1(parsed)
	log.Printf("Part 1 answer: %d\n", result)
}

func (s Solution) Part2(input []string) {
	parsed := input
	result := part2(parsed)
	log.Printf("Part 2 answer: %d\n", result)
}

func (s Solution) CheckPart1(input []string) error {
	parsed := input
	result := part1(parsed)
	log.Printf("Checking Part 1: got %d, expected %d\n", result, expPart1)
	if result != expPart1 {
		return fmt.Errorf("expected %d, got %d", expPart1, result)
	}
	return nil
}

func (s Solution) CheckPart2(input []string) error {
	parsed := input
	result := part2(parsed)
	log.Printf("Checking Part 2: got %d, expected %d\n", result, expPart2)
	if expPart2 == -9999 {
		log.Printf("No expected value for Part 2, skipping check")
		return nil
	}
	if result != expPart2 {
		return fmt.Errorf("expected %d, got %d", expPart2, result)
	}
	return nil
}
