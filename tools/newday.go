// go run tools/newday.go -year 2020 -day 1 [-exp1 0] [-exp2 0] [-force]
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func must(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func writeFile(path, content string, force bool) error {
	if !force {
		if _, err := os.Stat(path); err == nil {
			return fmt.Errorf("exists: %s (use -force to overwrite)", path)
		}
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0o644)
}

func main() {
	var (
		year  = flag.Int("year", 0, "AoC year, e.g. 2020")
		day   = flag.Int("day", 0, "AoC day, 1..25")
		exp1  = flag.Int("exp1", 0, "expected Part1 for sample or check (0 = none)")
		force = flag.Bool("force", false, "overwrite existing files")
	)
	flag.Parse()
	if *year == 0 || *day <= 0 || *day > 25 {
		panic("usage: -year YYYY -day N [ -exp1 X -force ]")
	}
	DD := fmt.Sprintf("%02d", *day)

	mainGo := fmt.Sprintf(`package main

import (
	"log"
	"os"

	"github.com/insomnes/goaoc/internal/run"
)

const (
	year = %d
	day  = %d
)

func main() {
	// Check If anything is provided, then run in solve mode
	if len(os.Args) > 1 {
		if err := run.Solve(year, day, Solution{}); err != nil {
			log.Fatalf("Error: %%v", err)
		}
	} else {
		if err := run.Check(year, day, Solution{}); err != nil {
			log.Fatalf("Error: %%v", err)
		}
	}
}
`, *year, *day)

	// Use your provided Solution, with exp values injected.
	// Note: 'toMeet' is specific to 2020 day 1; change as needed per puzzle.
	exp1Lit := " = -1 // No expected value provided yet"
	if *exp1 != 0 {
		exp1Lit = fmt.Sprintf(" = %d", *exp1)
	}
	exp2Lit := " = -1 // No expected value provided yet"

	solutionGo := fmt.Sprintf(`package main

import (
	"fmt"
	"log"

	"github.com/insomnes/goaoc/internal/measure"
	"github.com/insomnes/goaoc/internal/parse"
)

const (
	expPart1%s
	expPart2%s
)

type ParsedInput = []int


func part1(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 1")()
	log.Fatalf("PART 1: No solution found")
	return -1
}

func part2(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 2")()
	log.Fatalf("PART 2: No solution found")
	return -1
}

type Solution struct{}

func (s Solution) Part1(input []string) {
	parsed := parse.ToIntegers(input)
	result := part1(parsed)
	log.Printf("Part 1 answer: %%d\n", result)
}

func (s Solution) Part2(input []string) {
	parsed := parse.ToIntegers(input)
	result := part2(parsed)
	log.Printf("Part 2 answer: %%d\n", result)
}

func (s Solution) CheckPart1(input []string) error {
	parsed := parse.ToIntegers(input)
	result := part1(parsed)
	log.Printf("Checking Part 1: got %%d, expected %%d\n", result, expPart1)
	if result != expPart1 {
		return fmt.Errorf("expected %%d, got %%d", expPart1, result)
	}
	return nil
}

func (s Solution) CheckPart2(input []string) error {
	parsed := parse.ToIntegers(input)
	result := part2(parsed)
	log.Printf("Checking Part 2: got %%d, expected %%d\n", result, expPart2)
	if result != expPart2 {
		return fmt.Errorf("expected %%d, got %%d", expPart2, result)
	}
	return nil
}
`, exp1Lit, exp2Lit)

	inputFile := ""
	sampleFile := ""

	// Paths
	baseCmd := filepath.Join("cmd", fmt.Sprintf("%d", *year), DD)
	baseInputs := filepath.Join("inputs", fmt.Sprintf("y_%d", *year))
	fpMain := filepath.Join(baseCmd, "main.go")
	fpSol := filepath.Join(baseCmd, "solution.go")
	fpIn := filepath.Join(baseInputs, fmt.Sprintf("d_%s.txt", DD))
	fpSample := filepath.Join(baseInputs, fmt.Sprintf("d_%s_sample.txt", DD))

	// Write files
	log.Printf("Creating files for Year %d Day %02d", *year, *day)
	log.Printf("main: %s", fpMain)
	must(writeFile(fpMain, mainGo, *force))
	log.Printf("solution: %s", fpSol)
	must(writeFile(fpSol, solutionGo, *force))
	log.Printf("input: %s", fpIn)
	must(writeFile(fpIn, inputFile, *force))
	log.Printf("sample: %s", fpSample)
	must(writeFile(fpSample, sampleFile, *force))

	fmt.Printf("Created:\n  %s\n  %s\n  %s\n  %s\n", fpMain, fpSol, fpIn, fpSample)
}
