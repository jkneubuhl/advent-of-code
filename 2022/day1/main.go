package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "sort"
    "flag"
)

var input = flag.String("input", "input", "Input file")

func main() {

    flag.Parse()

    // read input
    file, err := os.Open(*input)
    if err != nil {
        fmt.Errorf("error %v", err)
    }
    defer file.Close()


    elves := []int{}
    total := 0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()

        cals, _ := strconv.Atoi(line)

        if cals == 0 {
            elves = append(elves, total)
            total = 0
            continue
        }

        total += cals
    }

    elves = append(elves, total)

    sort.Ints(elves)

    fmt.Printf("fattest of %d elves:\n", len(elves))

    topthree := 0
    for _, elf := range elves[len(elves)-3:] {
        fmt.Printf("%d\n", elf)

        topthree += elf
    }

    fmt.Printf("total: %d\n", topthree)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}