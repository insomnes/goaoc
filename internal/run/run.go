package run

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const inputBasePath = "inputs"

func makeInputPath(year int, day int, isSample bool) string {
	var filename string
	if isSample {
		filename = fmt.Sprintf("y_%d/d_%02d_sample.txt", year, day)
	} else {
		filename = fmt.Sprintf("y_%d/d_%02d.txt", year, day)
	}
	return filepath.Join(inputBasePath, filename)
}

func LoadInput(fp string) ([]string, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

type Solver interface {
	Part1(input []string)
	Part2(input []string)

	CheckPart1(input []string) error
	CheckPart2(input []string) error
}

func Solve(year int, day int, solver Solver) error {
	log.Printf("Solving Year %d Day %02d", year, day)
	inputStart := time.Now()

	inputPath := makeInputPath(year, day, false)
	input, err := LoadInput(inputPath)
	if err != nil {
		return err
	}
	inputDuration := time.Since(inputStart)
	log.Printf("Loaded input in %s", inputDuration)

	solveStart := time.Now()
	solver.Part1(input)
	solveDuration := time.Since(solveStart)
	log.Printf("Full part 1 took %s", solveDuration)

	solveStart = time.Now()
	solver.Part2(input)
	solveDuration = time.Since(solveStart)
	log.Printf("Full part 2 took %s", solveDuration)
	return nil
}

func Check(year int, day int, solver Solver) error {
	log.Printf("Checking Year %d Day %02d", year, day)
	inputPath := makeInputPath(year, day, true)
	input, err := LoadInput(inputPath)
	if err != nil {
		return err
	}

	if err := solver.CheckPart1(input); err != nil {
		return fmt.Errorf("part 1 check failed: %w", err)
	}
	log.Printf("Part 1 check passed")

	if err := solver.CheckPart2(input); err != nil {
		return fmt.Errorf("part 2 check failed: %w", err)
	}
	log.Printf("Part 2 check passed")
	return nil
}
