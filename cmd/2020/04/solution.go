package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/insomnes/goaoc/internal/measure"
	"github.com/insomnes/goaoc/internal/parse"
)

const (
	expPart1 = 2
	expPart2 = 1 // Change first byr to 1919 to make it invalid and get 1 here
)

// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID)
//

type Passport struct {
	BirthYear      uint16
	IssueYear      uint16
	ExpirationYear uint16
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      uint16
}

var BadPassport Passport

func parsePassport(data []string) Passport {
	p := Passport{}
	for _, line := range data {
		// Parse each line and populate the Passport struct
		for field := range strings.SplitSeq(line, " ") {
			kv := strings.SplitN(field, ":", 2)
			key, value := kv[0], kv[1]
			switch key {
			case "byr":
				p.BirthYear = parse.MustBeUint[uint16](value)
			case "iyr":
				p.IssueYear = parse.MustBeUint[uint16](value)
			case "eyr":
				p.ExpirationYear = parse.MustBeUint[uint16](value)
			case "hgt":
				p.Height = value
			case "hcl":
				p.HairColor = value
			case "ecl":
				p.EyeColor = value
			case "pid":
				p.PassportID = value
			case "cid":
				p.CountryID = parse.MustBeUint[uint16](value)
			}
		}
	}

	return p
}

func (p *Passport) ALlNeededFieldsPresent() bool {
	return p.BirthYear != BadPassport.BirthYear &&
		p.IssueYear != BadPassport.IssueYear &&
		p.ExpirationYear != BadPassport.ExpirationYear &&
		p.Height != BadPassport.Height &&
		p.HairColor != BadPassport.HairColor &&
		p.EyeColor != BadPassport.EyeColor &&
		p.PassportID != BadPassport.PassportID
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
// hgt (Height) - a number followed by either cm or in:
//
//	If cm, the number must be at least 150 and at most 193.
//	If in, the number must be at least 59 and at most 76.
//
// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
// pid (Passport ID) - a nine-digit number, including leading zeroes.
// cid (Country ID) - ignored, missing or not.
func (p *Passport) BirthYearValid() bool {
	return p.BirthYear >= 1920 && p.BirthYear <= 2002
}

func (p *Passport) IssueYearValid() bool {
	return p.IssueYear >= 2010 && p.IssueYear <= 2020
}

func (p *Passport) ExpirationYearValid() bool {
	return p.ExpirationYear >= 2020 && p.ExpirationYear <= 2030
}

func (p *Passport) HeightValid() bool {
	if strings.HasSuffix(p.Height, "cm") {
		value := parse.MustBeUint[uint16](strings.TrimSuffix(p.Height, "cm"))
		return value >= 150 && value <= 193
	} else if strings.HasSuffix(p.Height, "in") {
		value := parse.MustBeUint[uint16](strings.TrimSuffix(p.Height, "in"))
		return value >= 59 && value <= 76
	}
	return false
}

func (p *Passport) HairColorValid() bool {
	if len(p.HairColor) != 7 || p.HairColor[0] != '#' {
		return false
	}
	for _, ch := range p.HairColor[1:] {
		if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f')) {
			return false
		}
	}
	return true
}

func (p *Passport) EyeColorValid() bool {
	validColors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	return validColors[p.EyeColor]
}

func (p *Passport) PassportIDValid() bool {
	if len(p.PassportID) != 9 {
		return false
	}
	for _, ch := range p.PassportID {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

func (p *Passport) IsValid() bool {
	return p.ALlNeededFieldsPresent() &&
		p.BirthYearValid() &&
		p.IssueYearValid() &&
		p.ExpirationYearValid() &&
		p.HeightValid() &&
		p.HairColorValid() &&
		p.EyeColorValid() &&
		p.PassportIDValid()
}

type ParsedInput = []Passport

func parseInput(input []string) ParsedInput {
	defer measure.ExecutionTimeOfParsing("input", len(input))()
	passports := make([]Passport, 0)
	var currentPassportData []string
	for _, line := range input {
		if line == "" {
			// End of current passport data
			if len(currentPassportData) == 0 {
				log.Fatalf("Empty passport data encountered")
			}
			passports = append(passports, parsePassport(currentPassportData))
			currentPassportData = nil
			continue
		}
		currentPassportData = append(currentPassportData, line)
	}
	// Add the last passport if any
	if len(currentPassportData) > 0 {
		passports = append(passports, parsePassport(currentPassportData))
	}
	log.Printf("Parsed %d passports", len(passports))
	return passports
}

func part1(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 1")()
	validCount := 0
	for _, passport := range input {
		if passport.ALlNeededFieldsPresent() {
			validCount++
		}
	}
	return validCount
}

func part2(input ParsedInput) int {
	defer measure.ExecutionTimeOf("Part 2")()
	validCount := 0
	for _, passport := range input {
		if passport.IsValid() {
			validCount++
		}
	}
	return validCount
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
