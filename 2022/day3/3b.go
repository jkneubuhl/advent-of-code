package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var input = flag.String("in", "day3/input", "Input file")

type Squad [3]string

func main() {
	fmt.Println("--- Day 3: Rucksack Ruckus ---")

	flag.Parse()

	squads, err := load(*input)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, squad := range squads {
		fmt.Printf("%-32s\t", squad)

		badge := intersect(squad)
		priority := priority(badge)

		fmt.Printf("\t%d (%c)\n", priority, badge)

		total += priority
	}

	fmt.Printf("======================================\n")
	fmt.Printf("Reorganized %d squads.\n", len(squads))
	fmt.Printf("Badge priority sum: %d\n", total)
}

func priority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item-'A') + 1 + 26

	} else {
		return int(item-'a') + 1
	}
}

func intersect(squad Squad) rune {
	// map intersection would be better
	for _, r := range squad[0] {
		if strings.ContainsRune(squad[1], r) && strings.ContainsRune(squad[2], r) {
			return r
		}
	}

	return '?'
}

func load(file string) ([]Squad, error) {

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	squads := []Squad{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		sack1 := scanner.Text()

		scanner.Scan()
		sack2 := scanner.Text()

		scanner.Scan()
		sack3 := scanner.Text()

		squad := [3]string{sack1, sack2, sack3}
		squads = append(squads, squad)
	}

	return squads, nil
}
