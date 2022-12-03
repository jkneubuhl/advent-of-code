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

type Rucksack struct {
	Top, Bottom string
}

func main() {
	fmt.Println("--- Day 3: Rucksack Ruckus ---")

	flag.Parse()

	sacks, err := load(*input)
	if err != nil {
		panic(err)
	}

	total := 0
	for _, sack := range sacks {
		fmt.Printf("%-24s\t%-24s", sack.Top, sack.Bottom)

		common := intersect(sack.Top, sack.Bottom)
		priority := priority(common)

		fmt.Printf("\t%d (%c)\n", priority, common)

		total += priority
	}

	fmt.Printf("======================================\n")
	fmt.Printf("Reorganized %d rucksacks.\n", len(sacks))
	fmt.Printf("Mismatch sum: %d\n", total)
}

func priority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item-'A') + 1 + 26

	} else {
		return int(item-'a') + 1
	}
}

func intersect(c1, c2 string) rune {
	for _, r := range c1 {
		if strings.ContainsRune(c2, r) {
			return r
		}
	}

	return '?'
}

func load(file string) ([]*Rucksack, error) {
	sacks := []*Rucksack{}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		top := string(runes[:len(runes)/2])
		bottom := string(runes[len(runes)/2:])

		sack := &Rucksack{
			Top:    top,
			Bottom: bottom,
		}

		sacks = append(sacks, sack)
	}

	return sacks, nil
}
