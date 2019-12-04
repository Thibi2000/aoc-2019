package main

import (
	"fmt"
	"math"
)

type Pw struct {
	numbers []int
	i_end   int // index of last number, this is biting me in the ass tbh
}

func NewPw(length int) Pw {
	res := Pw{make([]int, length, length), 0}
	return res
}

func IntToPw(number int) Pw {
	numbers := []int{}
	for i := 1; number != 0; i++ {
		numbers = append(numbers, number%10)
		number /= 10
	}
	l := len(numbers)
	res := NewPw(l)
	for i := l - 1; i >= 0; i-- {
		res.numbers[i] = numbers[l-1-i]
	}
	res.i_end = l - 1
	return res
}

func Increment(pw *Pw) {
	current := &pw.numbers[pw.i_end]
	if *current++; *current == 10 {
		for i := pw.i_end; i >= 0 && pw.numbers[i] == 10; i-- {
			if i != 0 {
				pw.numbers[i] = 0
				pw.numbers[i-1]++
			} else {
				pw.numbers[i] = 1
				pw.i_end++
			}
		}
	}
}

func PwToInt(pw *Pw) int {
	l := len(pw.numbers)
	var res int = pw.numbers[l-1]
	for i := l - 2; i >= 0; i-- {
		if pw.numbers[i] != 0 {
			res += int(math.Pow(10, float64(l-1-i))) * pw.numbers[i]
		}
	}
	return int(res)
}

func findCorrectPw(pw *Pw, max int) (correct bool, max_reached bool) {
	equals := 0
	prev := pw.numbers[0]
	for i := 1; i < len(pw.numbers); i++ {
		//for prev > pw.numbers[i] {
        //   Increment(pw) // this iterates over everything and is not necessary
		//}
        if prev > pw.numbers[i] {
            for j:= i; j < len(pw.numbers); j++ { 
                pw.numbers[j] = prev
            }
            return true, PwToInt(pw) >= max
        } 
		if prev == pw.numbers[i] {
			equals++
		}
		prev = pw.numbers[i]
	}
    correct = equals != 0
	max_reached = PwToInt(pw) >= max
    return correct, max_reached
}

func hasDoubles(pw *Pw) bool {
    doubles := [10]int {}
    for i := 0; i < len(pw.numbers); i++ {
        doubles[pw.numbers[i]]++
    }
    for i := 0; i < 10; i++ {
        if doubles[i] == 2 {
            return true
        }
    }
    return false
}

func main() {
	pw := IntToPw(171309)
	Increment(&pw)
	var amt int
    // for part one
    // s/_/correct/ and hasDoubles -> correct
    _, max_reached := findCorrectPw(&pw, 643603)
	for !max_reached {
		if  hasDoubles(&pw) {
			amt++
		}
        Increment(&pw)
        _, max_reached = findCorrectPw(&pw, 643603)
	}
	fmt.Printf("Amt: %d\n", amt)
}
