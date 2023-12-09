package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// This program will probably never actually terminate in your PC. I just ran it for a while, collected the output, killed the program and then did math with the output. Sorry!
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
	var paths []string
	var path_starts []string
	for scanner.Scan() {
		identifiers := r_identifiers.FindAllString(scanner.Text(), -1)
		desert_map[identifiers[0]] = [2]string{identifiers[1], identifiers[2]}
		if identifiers[0][2] == 'A' {
			paths = append(paths, identifiers[0])
			path_starts = append(path_starts, identifiers[0])
		}
	}
	var paths_finished int
	var count int
	for idx := 0; paths_finished != len(paths); idx, count = (idx+1)%len(instructions), count+1 {
		current_instruction := instructions[idx]
		paths_finished = 0
		//fmt.Printf("Starting %dth step, going %c \n", count, current_instruction)
		for i := 0; i < len(paths); i++ {
			if current_instruction == 'L' {
				paths[i] = desert_map[paths[i]][0]
			} else {
				paths[i] = desert_map[paths[i]][1]
			}
			if paths[i][2] == 'Z' {
				fmt.Printf("\tFound %v (%v, %v). on path %d, step %d, starting point %v\n", paths[i], desert_map[paths[i]][0], desert_map[paths[i]][1], i, count, path_starts[i])
				paths_finished++
			}
		}
	}
	fmt.Printf("Reached the following nodes: \n")
	for i := 0; i < len(paths); i++ {
		fmt.Printf("\t%v\n", paths[i])
	}
	fmt.Printf("This took %d steps. \n", count)
}
