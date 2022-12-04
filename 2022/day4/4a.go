package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var input = flag.String("in", "day4/input", "Input file")

var cheerColumns = 0

type Range struct {
	Min, Max int
}

type Pair struct {
	Left, Right Range
}

func main() {
	fmt.Printf("--- Day 4: Camp Cleanup ---\n")

	flag.Parse()

	pairs, err := load(*input)
	if err != nil {
		panic(err)
	}

	overlaps := 0
	for _, pair := range pairs {
	    cheeryPrint(pair)

		if contains(pair.Left, pair.Right) || contains(pair.Right, pair.Left) {
			fmt.Printf("\t✅\n")
			overlaps++

        } else {
    	    fmt.Printf("\t❌\n")
        }

		fmt.Println()
	}

	fmt.Printf("======================================\n")
	fmt.Printf("Total pairs:       %d\n", len(pairs))
	fmt.Printf("Overlapping pairs: %d\n", overlaps)
}

// Does outer fully contain an inner range?
func contains(outer, inner Range) bool {
	return inner.Min >= outer.Min && inner.Max <= outer.Max
}

func cheeryPrint(pair Pair) {
    print(pair.Left, "⭐")
    fmt.Println()
    print(pair.Right, "🎄")
}

func print(r Range, icon string) {
    for i := 1; i <= cheerColumns; i++ {
        spec := Range{
            Min: i,
            Max: i,
        }

        if contains(r, spec) {
            fmt.Printf(icon)

        } else {
            fmt.Printf("➖")
        }
    }

    fmt.Printf("%v", r)
}

func load(file string) ([]Pair, error) {

	pairs := []Pair{}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		pair := Pair{}

		fmt.Sscanf(scanner.Text(),
			"%d-%d,%d-%d",
			&pair.Left.Min, &pair.Left.Max,
			&pair.Right.Min, &pair.Right.Max)

        // cheery columns
		if pair.Left.Max > cheerColumns {
		    cheerColumns = pair.Left.Max
		}

		if pair.Right.Max > cheerColumns {
		    cheerColumns = pair.Right.Max
		}

		pairs = append(pairs, pair)
	}

	return pairs, nil
}
