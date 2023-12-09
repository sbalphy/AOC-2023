package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	desert_map := make(map[string][2]string)
	// back to da old me
	r_identifiers, _ := regexp.Compile("[A-Z][A-Z][A-Z]")
	scanner.Scan()
	instructions := scanner.Text()
	// come o newline
	scanner.Scan()
	for scanner.Scan() {
		identifiers := r_identifiers.FindAllString(scanner.Text(), -1)
		desert_map[identifiers[0]] = [2]string{identifiers[1], identifiers[2]}
	}
	current_node := "AAA"
	var count int
	for idx := 0; current_node[2] != 'Z'; idx, count = (idx+1)%len(instructions), count+1 {
		fmt.Printf("Currently at %v (%v, %v), going %c \n", current_node, desert_map[current_node][0], desert_map[current_node][1], instructions[idx])
		if instructions[idx] == 'L' {
			current_node = desert_map[current_node][0]
		} else {
			current_node = desert_map[current_node][1]
		}
	}
	fmt.Printf("Reached node %v after %d steps. \n", current_node, count)
}
