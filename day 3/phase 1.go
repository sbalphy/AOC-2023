package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ispart(line int, start int, end int, schematic []string) bool {
	// definimos uma bounding box em torno do numero e checamos todos os caracteres dentro dela
	bounds_y_start := max(0, line-1)
	bounds_y_end := min(len(schematic)-1, line+1)
	bounds_x_start := max(0, start-1)
	bounds_x_end := min(len(schematic[0])-1, end+1)
	for i := bounds_y_start; i <= bounds_y_end; i++ {
		for j := bounds_x_start; j <= bounds_x_end; j++ {
			if issymbol(schematic[i][j]) {
				return true
			}
		}
	}
	return false
}

func isnumeric(char byte) bool {
	return char >= 48 && char <= 57
}

func issymbol(char byte) bool {
	return (char < 48 || char > 57) && char != 46
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	r_numbers, _ := regexp.Compile("[0-9]+")
	var schematic []string
	var sum int
	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}
	for i := 0; i < len(schematic); i++ {
		numbers_idx := r_numbers.FindAllStringIndex(schematic[i], -1)
		for j := 0; j < len(numbers_idx); j++ {
			start := numbers_idx[j][0]
			end := numbers_idx[j][1] - 1
			value, _ := strconv.Atoi(schematic[i][start : end+1])
			part := ispart(i, start, end, schematic)
			if part {
				sum += value
			}
			fmt.Printf("Found number %d in line %d. Part = %t. Current sum: %d \n", value, i, part, sum)
		}
	}
	fmt.Printf("final sum %d \n", sum)
}
