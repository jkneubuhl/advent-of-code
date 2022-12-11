package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var input = flag.String("in", "day8/input", "Input file")

type Tree struct {
	Visible bool
	Height  int
}

func main() {

	forest := load()

	count := scan(forest)

	print(forest)

	fmt.Printf("Total visible: %d\n", count)
}

func scan(forest [][]*Tree) int {

	// l->r scan
	for row := 0; row < len(forest); row++ {
		max := -1
		for col := 0; col < len(forest[row]); col++ {
			if forest[row][col].Height > max {
				forest[row][col].Visible = true
				max = forest[row][col].Height
			}
		}
	}

	// r->l scan
	for row := 0; row < len(forest); row++ {
		max := -1
		for col := len(forest[row]) - 1; col >= 0; col-- {
			if forest[row][col].Height > max {
				forest[row][col].Visible = true
				max = forest[row][col].Height
			}
		}
	}

	// t->b scan
	for col := 0; col < len(forest[0]); col++ {
		max := -1
		for row := 0; row < len(forest); row++ {
			if forest[row][col].Height > max {
				forest[row][col].Visible = true
				max = forest[row][col].Height
			}
		}
	}

	// b->t scan
	for col := len(forest[0]) - 1; col >= 0; col-- {
		max := -1
		for row := len(forest) - 1; row >= 0; row-- {
			if forest[row][col].Height > max {
				forest[row][col].Visible = true
				max = forest[row][col].Height
			}
		}
	}

	// count visible trees
	visible := 0
	for _, hedge := range forest {
		for _, tree := range hedge {
			if tree.Visible {
				visible++
			}
		}
	}

	return visible
}

func print(forest [][]*Tree) {

	for _, hedge := range forest {
		for _, tree := range hedge {
			if tree.Visible {
				//fmt.Printf("%v ", tree)
				fmt.Printf("🎄")

			} else {
				fmt.Printf("➖")
			}
		}
		fmt.Println()
	}
}

func load() [][]*Tree {

	flag.Parse()

	forest := [][]*Tree{}

	f, err := os.Open(*input)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		//		fmt.Println(line)

		hedge := []*Tree{}
		for _, height := range line {
			h, _ := strconv.Atoi(string(height))
			tree := &Tree{
				Visible: false,
				Height:  h,
			}

			hedge = append(hedge, tree)
		}

		forest = append(forest, hedge)
		row++
	}

	return forest
}
