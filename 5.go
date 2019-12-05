package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var test_inputs = [][]int{
	{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
	{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
	{3, 3, 1108, -1, 8, 3, 4, 3, 99},
	{3, 3, 1107, -1, 8, 3, 4, 3, 99},
    {3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9},
	{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
    {3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,
        1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,
        999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99},
}

type Instruction struct {
	opcode      int
	args_length int
	do          func(args []int) (res int)
}

var instructions = []Instruction{
	Instruction{1, 3, func(args []int) int {
		return args[0] + args[1]
	}},
	Instruction{2, 3, func(args []int) int {
		return args[0] * args[1]
	}},
	Instruction{3, 1, func(args []int) int {
		var i int
		fmt.Scanf("%d", &i)
		return i
	}},
	Instruction{4, 1, func(args []int) int {
		fmt.Println(args[0])
		return args[0]
	}},
	Instruction{5, 2, func(args []int) int {
		if args[0] != 0 {
			return args[1]
		}
		return -1 // you can return negative numbers because they're indexes
	}},
	Instruction{6, 2, func(args []int) int {
		if args[0] == 0 {
			return args[1]
		}
		return -1
	}},
	Instruction{7, 3, func(args []int) int {
		if args[0] < args[1] {
			return 1
		}
		return 0
	}},
	Instruction{8, 3, func(args []int) int {
		if args[0] == args[1] {
			return 1
		}
		return 0
	}},
	// not sure is I should add Instruction{99, 0, func(args []int)(int, int) {return 0, -1}}, // -1 so we can check to halt
}

type Word struct {
	opcode int
	modes  []int
}

func NewWord(ins int) Word {
    fmt.Print(ins)
	opcode := ins % 10
	ins /= 10
	opcode += 10 * (ins % 10)
	ins /= 10
	modes := make([]int, 3, 3)
	for i := 0; ins != 0; i++ {
		modes[i] = ins % 10
		ins /= 10
	}
	res := Word{opcode, modes}
	fmt.Println("\t", res)
	return res
}
func Exec(opcodes []int) {
	size := len(opcodes)
	pos := 0
	ins := opcodes[pos]
	word := NewWord(ins)
	for word.opcode != 99 {
		instruction := instructions[word.opcode-1]
		if pos+instruction.args_length > size {
			break
		}
		args := make([]int, instruction.args_length, instruction.args_length)
		for i, el := range opcodes[pos+1 : pos+instruction.args_length+1] {
			if word.modes[i] == 0 {
				args[i] = opcodes[el]
			} else {
				args[i] = el
			}
		}
        res := instruction.do(args)
        // ugly 
        if word.opcode == 5 ||
            word.opcode == 6 {
            if res != -1 {
                pos = res
            } else {
                pos += instruction.args_length + 1
            }
        } else {
            resPos := opcodes[pos+instruction.args_length]
            if word.opcode != 4 {
                opcodes[resPos] = res
            }
            pos += instruction.args_length + 1
        }
        word = NewWord(opcodes[pos])
	}
}

func Part1(memory []int, size int) int {
	input := make([]int, size, size)
	copy(input, memory)
	input[1] = 12
	input[2] = 2
	Exec(input)
	return input[0]
}

func Part2(memory []int, size int) (int, int) {
	input := make([]int, size, size)
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			copy(input, memory)
			input[1] = noun
			input[2] = verb
			Exec(input)
			if input[0] == 19690720 {
				return noun, verb
			}
		}
	}
	return 0, 0
}

func main() {
	file, _ := os.Open("inputs/5.txt") // no error checking, assume it's there
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan() // there's only one line in the file
	s := strings.Split(scanner.Text(), ",")
	size := len(s)
	memory := make([]int, size, size)
	for pos, el := range s {
		memory[pos], _ = strconv.Atoi(el)
	}
    //for _, input := range test_inputs {
    //Exec(test_inputs[6])
        Exec(memory)
    //}
}
