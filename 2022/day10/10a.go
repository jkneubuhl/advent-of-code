package main

import (
    "bufio"
    "os"
    "fmt"
    "flag"
)

var input = flag.String("in", "day10/input", "Input file")

const maxCycle = 220

type Instruction struct {
    Operator string
    Operand int
    CycleCount int
}

func main() {
    fmt.Printf("--- Day 10: Cathode-Ray Tube\n")

    instructions := load()

    regx := 1
    pc := 0
    total := 0

    for cycle := 1; cycle <= maxCycle && pc < len(instructions); cycle++ {

        // start of cycle
        in := instructions[pc]
        fmt.Printf("START\t%d\t%d\t%v\n", cycle, regx, in)

        strength := cycle * regx

        if (cycle-20) % 40 == 0 {
            total += strength
            fmt.Printf("** strength: %d\n", strength)
        }

        in.CycleCount--

        // end of cycle
        if in.CycleCount == 0 {
            if in.Operator == "addx" {
                regx += in.Operand
            }

            // next
            pc++
        }

        fmt.Printf("END\t%d\t%d\t%v\n", cycle, regx, in)
    }

    fmt.Printf("================================================================================\n")
    fmt.Printf("total: %d\n", total)
}

func load() []*Instruction {
    flag.Parse()

    f, err := os.Open(*input)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    instructions := []*Instruction{}

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()

        in := &Instruction{}

        fmt.Sscanf(line, "%s %d", &in.Operator, &in.Operand)

        if in.Operator == "noop" {
            in.CycleCount = 1
        }

        if in.Operator == "addx" {
            in.CycleCount = 2
        }

        instructions = append(instructions, in)
    }

    return instructions
}