package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var input = flag.String("in", "day2/input", "Input file")

type Play string

type Outcome string

const (
	Rock    Play = "A"
	Paper        = "B"
	Scissor      = "C"

	Lose Outcome = "X"
	Draw         = "Y"
	Win          = "Z"
)

type Round struct {
	Elf  Play
	Goal Outcome
	My   Play
}

func main() {
	fmt.Println("Rock Paper Elf")

	flag.Parse()

	game, err := loadGame(*input)
	if err != nil {
		panic(err)
	}

	total := 0
	for count, round := range game {
		chooseMove(round)

		shape   := scoreShape(round.My)
		outcome := scoreOutcome(round)
		score   := shape + outcome

		fmt.Printf("round %d: %v -> %d + %d = %d\n", count, round, shape, outcome, score)

		total += score
	}

	fmt.Printf("Total rounds: %d\n", len(game))
	fmt.Printf("Total score:  %d\n", total)
}

func scoreShape(play Play) int {
	switch play {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissor:
		return 3
	}

	return 0
}

// diff modulo 3
//  0 : draw
//  1 : win
//  2 : lose
func outcome(round *Round) Outcome {
	diff := int(round.My[0]) - int(round.Elf[0])
	modulo := (diff%3 + 3) % 3

	switch modulo {
	case 0:
		return Draw
	case 1:
		return Win
	case 2:
		return Lose
	}

	panic("bad game")
	return Lose
}

func scoreOutcome(round *Round) int {
	switch outcome(round) {
	case Lose:
		return 0
	case Draw:
		return 3
	case Win:
		return 6
	}

	panic("bad outcome")
	return 0
}

func chooseMove(round *Round) {

	if round.Goal == Lose {
		switch round.Elf {
		case Rock:
			round.My = Scissor
		case Paper:
			round.My = Rock
		case Scissor:
			round.My = Paper
		}

		return
	}

	if round.Goal == Win {
		switch round.Elf {
		case Rock:
			round.My = Paper
		case Paper:
			round.My = Scissor
		case Scissor:
			round.My = Rock
		}

		return
	}

	if round.Goal == Draw {
		round.My = round.Elf

		return
	}
}

func loadGame(file string) ([]*Round, error) {

	game := []*Round{}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line   := scanner.Text()
		tokens := strings.Fields(line)
		round  := &Round{
			Elf:  parsePlay(tokens[0]),
			Goal: parseOutcome(tokens[1]),
		}

		game = append(game, round)
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return game, nil
}

func parsePlay(token string) Play {
	switch token {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissor
	}
	panic("Unknown token: " + token)
	return Rock
}

func parseOutcome(token string) Outcome {
	switch token {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	}

	panic("Unknown token: " + token)
	return Lose
}