package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var input = flag.String("in", "day6/input", "Input file")

func main() {
	fmt.Printf("--- Day 6: Tuning Trouble ---\n")

	flag.Parse()

	datastream, err := load(*input)
	if err != nil {
		panic(err)
	}

	width := 4 // 6a
	//	width := 14     // 6b
	for i := width; i <= len(datastream); i++ {
		buf := datastream[i-width : i]
		count := unique(buf)

		fmt.Printf("%d %s -> %d\n", i, buf, count)

		if count == width {
			fmt.Printf("first marker after character %d\n", i)
			break
		}
	}
}

// accurate but not optimal: count in map / bag
func unique(b string) int {
	m := map[rune]int{}
	for _, c := range b {
		m[c]++
	}

	return len(m)
}

func load(filename string) (string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	return scanner.Text(), nil
}
