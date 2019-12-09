package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type OrbitNode struct {
	value         string
	orbits_around *OrbitNode
}

func InitMap(filename string) map[string]*OrbitNode {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nodes := make(map[string]*OrbitNode)
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), ")")
		for _, el := range elements {
			if _, ok := nodes[el]; !ok {
				node := OrbitNode{el, nil}
				nodes[el] = &node
			}
		}
		nodes[elements[1]].orbits_around = nodes[elements[0]]
	}
	return nodes
}

func Part1() {
	nodes := InitMap("inputs/6.txt")
	orbits := 0
	for _, value := range nodes {
		for node := value; node.orbits_around != nil; orbits++ {
			node = node.orbits_around
		}
	}
	fmt.Println("Part1: ", orbits)

}

// look for common orbit
// sum of both path's to that orbit
func Part2() {
	nodes := InitMap("inputs/6.txt")
	santa_node, _ := nodes["SAN"]
	node, _ := nodes["YOU"]
	// key: name
	// int: path length
	nodes_connected_with_you := make(map[string]int)
	nodes_connected_with_san := make(map[string]int)
	prev := -2
	for node.orbits_around != nil {
		_, ok := nodes_connected_with_you[node.value]
		if !ok {
			nodes_connected_with_you[node.value] = prev
		}
		nodes_connected_with_you[node.value]++
		prev = nodes_connected_with_you[node.value]
		node = node.orbits_around
	}
	prev = -2
	for santa_node.orbits_around != nil {
		_, ok := nodes_connected_with_san[santa_node.value]
		if !ok {
			nodes_connected_with_san[santa_node.value] = prev
		}
		nodes_connected_with_san[santa_node.value]++
		prev = nodes_connected_with_san[santa_node.value]
		santa_node = santa_node.orbits_around
	}
	length := len(nodes_connected_with_san)
	smallest := nodes_connected_with_san
	biggest := nodes_connected_with_you
	if length < len(nodes_connected_with_you) {
		smallest = nodes_connected_with_you
		biggest = nodes_connected_with_san
	}
	min := 241064 // answer of part 1
    common := "YOU"
	for key, value := range smallest {
		if path_length, ok := biggest[key]; ok {
			if min > path_length+value {
				min = path_length + value
                common = key
			}
		}
	}
	fmt.Println("Part2: ", min, common)
}
func main() {
	Part2()
}
