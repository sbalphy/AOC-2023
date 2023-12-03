package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func gear_check(line int, col int, schematic []string) int {
	// definimos uma bounding box em torno da marcha e procuramos números em volta dela
	// retorna a ratio (se nao for gear, ratio eh -1)
	bounds_y_start := max(0, line-1)
	bounds_y_end := min(len(schematic)-1, line+1)
	// smile
	r_numbers, _ := regexp.Compile("[0-9]+")
	var adjacent_numbers []int
	for i := bounds_y_start; i <= bounds_y_end; i++ {
		numbers_idx := r_numbers.FindAllStringIndex(schematic[i], -1)
		for j := 0; j < len(numbers_idx); j++ {
			start := numbers_idx[j][0]
			end := numbers_idx[j][1] - 1
			value, _ := strconv.Atoi(schematic[i][start : end+1])
			if col <= end+1 && col >= start-1 {
				adjacent_numbers = append(adjacent_numbers, value)
			}
		}
	}
	if len(adjacent_numbers) != 2 {
		fmt.Printf("Found gear in line %d, column %d, but it is adjacent to %d parts. \n", line, col, len(adjacent_numbers))
		return -1
	}
	return adjacent_numbers[0] * adjacent_numbers[1]
}

func issymbol(char byte) bool {
	return (char < 48 || char > 57) && char != 46
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	var schematic []string
	var sum int
	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}
	// se eu usar regex pra achar 1 caractere Deus me mata
	for i := 0; i < len(schematic); i++ {
		for j := 0; j < len(schematic[i]); j++ {
			if schematic[i][j] == 42 {
				gear_ratio := gear_check(i, j, schematic)
				if gear_ratio != -1 {
					sum += gear_ratio
				}
				fmt.Printf("Current sum: %d \n", sum)
			}
		}
	}
	fmt.Printf("final sum %d \n", sum)
}
