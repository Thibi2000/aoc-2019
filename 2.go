package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var test_input = []int{
	1, 9, 10, 3,
	2, 3, 11, 0,
	99,
	30, 40, 50,
}

type Instruction struct {
	opcode      int
	args_length int
	do func(args []int) (res int)
}

var instructions = []Instruction{
	Instruction{1, 3, func(args []int) int { return args[0] + args[1]}},
	Instruction{2, 3, func(args []int) int { return args[0] * args[1]}},
	// not sure is I should add Instruction{99, 0, func(args []int)(int, int) {return 0, -1}}, // -1 so we can check to halt
}

func Exec(opcodes []int) []int {
	size := len(opcodes)
	pos := 0
	opcode := opcodes[pos]
	for opcode != 99 {
		instruction := instructions[opcode-1]
		if pos+instruction.args_length > size {
			break
		}
		args := make([]int, instruction.args_length, instruction.args_length)
		for i, el := range opcodes[pos+1 : pos+instruction.args_length] {
			args[i] = opcodes[el]
		}
		resPos :=  opcodes[pos+instruction.args_length]
		res := instruction.do(args)
		opcodes[resPos] = res
		pos += instruction.args_length + 1
		opcode = opcodes[pos]
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

func Part2(memory []int, size int) (int, int) {
	input := make([]int, size, size)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(input, memory)
			input[1] = noun
			input[2] = verb
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
	fmt.Printf("Part2: %d\n", noun*100+verb)
}
