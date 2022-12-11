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

	max := 0
	for row := 0; row < len(forest); row++ {
		for col := 0; col < len(forest[row]); col++ {
			score := viewScore(forest, row, col)
			if score > max {
				max = score
			}
		}
	}

	fmt.Printf("max view score: %d\n", max)
}

func viewScore(forest [][]*Tree, row, col int) int {

	height := forest[row][col].Height

	var left, right, up, down int

	// -> left
	for j := col - 1; j >= 0; j-- {
		left++
		if forest[row][j].Height >= height {
			break
		}
	}

	// -> right
	for j := col + 1; j <= len(forest[col])-1; j++ {
		right++
		if forest[row][j].Height >= height {
			break
		}
	}

	// -> down
	for i := row + 1; i <= len(forest)-1; i++ {
		down++
		if forest[i][col].Height >= height {
			break
		}
	}

	// -> up
	for i := row - 1; i >= 0; i-- {
		up++
		if forest[i][col].Height >= height {
			break
		}
	}

	return left * right * up * down
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
