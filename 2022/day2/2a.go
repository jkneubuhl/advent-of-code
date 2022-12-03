package main

import (
    "os"
    "fmt"
    "bufio"
    "flag"
)

var input = flag.String("input", "input", "Input file")

type Hand int

type Round struct {
    human, elf Hand
}

const (
    Rock Hand = iota
    Paper
    Scissor
)

func loadGame(file string) ([]Round, error) {

    game := []Round{}

    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        elf := scanner.Text()
        scanner.Scan()
        human := scanner.Text()

        round := Round{
            elf:   parsePlay(elf),
            human: parsePlay(human),
        }

        game = append(game, round)
    }

    if err = scanner.Err(); err != nil {
        return nil, err
    }

    return game, nil
}

func parsePlay(token string) Hand {
    switch token {
        case "A":
            return Rock
        case "X":
            return Rock
        case "B":
            return Paper
        case "Y":
            return Paper
        case "C":
            return Scissor
        case "Z":
            return Scissor
    }
    panic("Unknown token: " + token)
    return Rock
}

func scoreShape(play Hand) int {
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

func scoreOutcome(round Round) int {

    diff := round.human - round.elf

    switch diff {
        // win
        case -2:
            return 6

        // loss
        case -1:
            return 0

        // draw
        case 0:
            return 3

        // win
        case 1:
            return 6

        // loss
        case 2:
            return 0
    }

    return 0
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
        shape := scoreShape(round.human)
        outcome := scoreOutcome(round)
        score := shape + outcome

        diff := round.human - round.elf

        fmt.Printf("round %d: %v -> %d -> (%d + %d = %d)\n", count, round, diff, shape, outcome, score)

        total += score
    }

    fmt.Printf("Total rounds: %d\n", len(game))
    fmt.Printf("Total score:  %d\n", total)
}
