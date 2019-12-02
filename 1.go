package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func DoEachLine(filename string, start_val int,  action func(line string)int ) int {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("Can't open file")
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    res := start_val
    for scanner.Scan() {
        res += action(scanner.Text())
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Println("Couldn't read file")
    }

    return res
}

func First() {
    part1 := DoEachLine("inputs/1.txt", 0, func(line string) int {
        mass, err := strconv.Atoi(line)
        if err != nil {
            fmt.Println("Couldn't convert text to int")
        }
        return mass / 3 - 2
    })
    part2 := DoEachLine("inputs/1.txt", 0, func(line string) int {
        mass, err := strconv.Atoi(line)
        if err != nil {
            fmt.Println("Couldn't convert text to int")
        }
        res := 0
        for f:= mass/3 - 2; f > 0; f = f / 3 - 2 {
            res += f
        }
        return res
    })
    fmt.Printf("Part1: %d\n", part1)
    fmt.Printf("Part2: %d\n", part2)
}

