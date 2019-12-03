package main

import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

type Instruction func(int, int)(int)
var instructions = [2]Instruction{
    func(a int, b int)(int) { return a + b},
    func(a int, b int)(int) { return a * b},
}

func Exec(opcodes []int) [] int {
    size := len(opcodes)
    for pos := 0; pos < size; pos += 4 {
        opcode := opcodes[pos]
        if opcode == 99 || pos + 3 > size {
            break;
        }
        nextPos := opcodes[pos + 3]
        opcodes[nextPos] = instructions[opcode - 1](opcodes[opcodes[pos + 1]], opcodes[opcodes[pos + 2]])
    }
    return opcodes
}
func Part1(memory []int, size int) int {
    input := make([]int, size, size)
    copy(input, memory)
    input[1] = 12
    input[2] = 2
    return Exec(input)[0]
}

func Part2(memory []int, size int) (int,int) {
    input := make([]int, size, size)
    for noun := 0; noun < 100; noun++ {
        for verb := 0; verb < 100; verb++ {
            copy(input, memory)
            input[1] = noun;
            input[2] = verb;
            if Exec(input)[0] == 19690720 {
                return noun, verb
            }
        }
    }
    return 0, 0
}

func main() {
    file, _ := os.Open("inputs/2.txt") // no error checking, assume it's there
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Scan() // there's only one line in the file
    s := strings.Split(scanner.Text(), ",")
    size := len(s)
    memory := make([]int, size, size)
    for pos, el := range s {
        memory[pos], _ = strconv.Atoi(el)
    }
    fmt.Printf("Part1: %d\n", Part1(memory, size))
    noun, verb := Part2(memory, size)
    fmt.Printf("Part2: %d %d\n", noun, verb)
}

