package main
// This seems messy and i had to work too long for this
// I'm nto that happy with the result
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var test_input = [][2]string{
	{"R8,U5,L5,D3",
		"U7,R6,D4,L4"},
	{"R75,D30,R83,U83,L12,D49,R71,U7,L72",
		"U62,R66,U55,R34,D71,R55,D58,R83"},
	{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
		"U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}}

// Trying the DOD meme
// Each line is represented with 2 indices
// So Line 1 is has index 0 as begin, and 1 as end point
type Lines struct {
	x []int
	y []int
}

func IsVertical(x1 int, x2 int) bool {
	return x1 == x2
}

func ManhattanDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
}

// checks if a is in between the 2 other numbers
// order doesn't matter
func Inbetween(a int, b int, c int) bool {
	max := b
	min := c
	if min > max {
		max = c
		min = b
	}
	return a <= max && a >= min
}

// if bool is false then there is no intersection
func Intersect(lines_a *Lines, index_a int, lines_b *Lines, index_b int) ([2]int, bool) {
	res := [2]int{}
	a_vert, b_vert := IsVertical(lines_a.x[index_a], lines_a.x[index_a+1]),
		IsVertical(lines_b.x[index_b], lines_b.x[index_b+1])
	if a_vert == b_vert {
		// they're parallel so there's no intersection
		return res, false
	} else {
		// I don't find this very pretty
		if a_vert &&
			Inbetween(lines_a.x[index_a], lines_b.x[index_b], lines_b.x[index_b+1]) &&
			Inbetween(lines_b.y[index_b], lines_a.y[index_a], lines_a.y[index_a+1]) {
			res[0] = lines_a.x[index_a]
			res[1] = lines_b.y[index_b]
		} else if b_vert &&
			Inbetween(lines_b.x[index_b], lines_a.x[index_a], lines_a.x[index_a+1]) &&
			Inbetween(lines_a.y[index_a], lines_b.y[index_b], lines_b.y[index_b+1]) {
			res[0] = lines_b.x[index_b]
			res[1] = lines_a.y[index_a]
		} else {
			return res, false
		}
		return res, true
	}
}

func InitLines(s string) Lines {
	steps := strings.Split(s, ",")
	length := len(steps)
	lines := Lines{}
	lines.x = make([]int, length+1, length+1) // +1 to include (0,0)
	lines.y = make([]int, length+1, length+1)
	lines.x[0] = 0
	lines.y[0] = 0
	next_x := 0
	next_y := 0
	for index, step := range steps {
		inc, _ := strconv.Atoi(string(step[1:]))
		switch step[0] {
		case 'R':
			next_x = lines.x[index] + inc
			next_y = lines.y[index]
		case 'L':
			next_x = lines.x[index] - inc
			next_y = lines.y[index]
		case 'U':
			next_x = lines.x[index]
			next_y = lines.y[index] + inc
		case 'D':
			next_x = lines.x[index]
			next_y = lines.y[index] - inc
		}
		lines.x[index+1] = next_x
		lines.y[index+1] = next_y
	}
	return lines
}

// Returns the point and the distance
func GetPoints(strings [2]string) [][2]int {
	lines := [2]Lines{}
	for i := 0; i < 2; i++ {
		lines[i] = InitLines(strings[i])
	}
	n_points1 := len(lines[0].x) // assume it will be the same for both lines
	n_points2 := len(lines[1].x) // assume it will be the same for both lines
	points := [][2]int{}
	// increment by 2 because the end of the line segment
	// has to become the beginning of the next
	for i := 0; i < n_points1-1; i++ {
		for j := 0; j < n_points2-1; j++ {
			if !(lines[0].x[i] == 0 && lines[0].y[i] == 0 &&
				lines[1].y[j] == 0 && lines[1].y[j] == 0) {
				if res, found := Intersect(&lines[0], i, &lines[1], j); found {
					points = append(points, res)
				}
			}
		}
	}
	return points
}

func part1(strings [2]string) {
    fmt.Println("PART1: ")
	points := GetPoints(strings)
	min_distance := ManhattanDistance(0, 0, points[0][0], points[0][1])
	min_index := 0
	for i := 0; i < len(points); i++ {
		dis := ManhattanDistance(0, 0, points[i][0], points[i][1])
		if dis <= min_distance {
			min_distance = dis
			min_index = i
		}
	}

	fmt.Printf("The point is %v with distance %d\n\n", points[min_index], min_distance)
}

func StepsToPoints(lines *Lines, points [][2]int) []int{
	steps := make([]int, len(points), len(points))
    step := 0
	for i := 0; i < len(lines.x)-1; i++ {
		//isVert := IsVertical(lines.x[i], lines.x[i+1])
		point_index := -1
        for index, p := range points {
            if Inbetween(p[1], lines.y[i], lines.y[i+1]) &&
                Inbetween(p[0], lines.x[i], lines.x[i+1]) {
                point_index = index
				break
			}
		}
		if point_index == -1 {
			step += ManhattanDistance(lines.x[i], lines.y[i], lines.x[i+1], lines.y[i+1])
		} else {
			step +=  ManhattanDistance(lines.x[i], lines.y[i], points[point_index][0], points[point_index][1])
            steps[point_index] = step
            step +=  ManhattanDistance(points[point_index][0], points[point_index][1], lines.x[i+1], lines.y[i+1]) 
		}
	}
    return steps
}

func part2(strings [2]string) {
    fmt.Println("PART 2:")
	points := GetPoints(strings)
	lines := [2]Lines{}
	for index, s := range strings {
		lines[index] = InitLines(s)
	}
    steps_line1 := StepsToPoints(&lines[0], points)
    steps_line2 := StepsToPoints(&lines[1], points)
    steps := []int{}
    for i := 0; i < len(points); i++ {
        if !(steps_line1[i] == 0 || steps_line2[i] == 0) {
            steps = append(steps, steps_line1[i] + steps_line2[i])
        }
    }
    min, index := min(steps)
    fmt.Printf("Min amount of steps: %d\nTo point: %v\n\n", min, points[index])
}

func min(l []int) (int, int) {
    min_index := 0
    min := l[min_index]
    for index, el := range l[1:] {
        if min >= el {
            min = el
            min_index = index + 1 // +1 because it starts at the second element
        }
    }
    return min, min_index
}
func main() {
	file, err := os.Open("inputs/3.txt")
	if err != nil {
		fmt.Println("Coudln't open file")
	}
	defer file.Close()
	strings := [2]string{}
	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan() && i < 2; i++ {
		strings[i] = scanner.Text()
	}
	part1(strings)
    part2(strings)
}
