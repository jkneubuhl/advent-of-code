package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var input = flag.String("in", "day5/input", "Input file")

type Stack struct {
	v []string
}

func (s *Stack) Push(c string) {
	s.v = append(s.v, c)
}

func (s *Stack) Top() string {
	return s.v[len(s.v)-1]
}

func (s *Stack) Pop() string {
	top := s.Top()
	s.v = s.v[:len(s.v)-1]

	return top
}

type Move struct {
	Count, From, To int
}

func main() {
	fmt.Printf("--- Day 5: Supply Stacks ---\n")

	flag.Parse()

	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	stacks := loadStacks(scanner)
	moves := loadMoves(scanner)

	cheeryPrint(stacks)

	for _, move := range moves {
		source := stacks[move.From-1]
		dest := stacks[move.To-1]

		// 5a : move directly from source to dest
		//         for i := 0; i < move.Count; i++ {
		//             dest.Push(source.Pop())
		//         }

		// 5b : retain order when moving the containers:
		temp := &Stack{}
		for i := 0; i < move.Count; i++ {
			temp.Push(source.Pop())
		}
		for i := 0; i < move.Count; i++ {
			dest.Push(temp.Pop())
		}

		fmt.Printf("after move %d from %d to %d:\n", move.Count, move.From, move.To)

		cheeryPrint(stacks)
	}

	fmt.Printf("======================================\n")
	fmt.Printf("Stack tops: ")
	for _, stack := range stacks {
		fmt.Printf("%s", stack.Top())
	}

	fmt.Println()
}

func loadStacks(scanner *bufio.Scanner) []*Stack {

	count := 0
	stacks := []*Stack{}

	// push the input lines onto a stack
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		// base line of the stack input.  How many?
		if strings.HasPrefix(line, " 1") {
			count = len(strings.Fields(line))
			break
		}

		lines = append(lines, line)
	}

	// consume the empty line separating the initial stacks from moves
	scanner.Scan()
	fmt.Printf("======================================\n")

	// start with empty stacks
	for i := 0; i < count; i++ {
		stacks = append(stacks, &Stack{})
	}

	// load up the stacks
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]

		// Parse each line of input by 4-char groups
		for pos := 0; pos < len(line); pos += 4 {
			token := line[pos : pos+3]

			if strings.TrimSpace(token) == "" {
				continue
			}

			// Push a container onto a stack.
			// *phew*
			container := string(token[1])

			stacknum := pos / 4
			stacks[stacknum].Push(container)
		}
	}

	return stacks
}

func cheeryPrint(stacks []*Stack) {
	for i, s := range stacks {
		fmt.Printf("%d %v\n", i+1, s)
	}

	fmt.Println()
}

func loadMoves(scanner *bufio.Scanner) []Move {
	moves := []Move{}

	for scanner.Scan() {
		line := scanner.Text()
		move := Move{}

		fmt.Sscanf(line,
			"move %d from %d to %d",
			&move.Count, &move.From, &move.To)

		moves = append(moves, move)
	}

	return moves
}
